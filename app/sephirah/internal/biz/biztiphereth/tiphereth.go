package biztiphereth

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type TipherethRepo interface {
	FetchUserByPassword(context.Context, *User) (*User, error)
	CreateUser(context.Context, *User, model.InternalID) error
	UpdateUser(context.Context, *User, string) error
	ListUser(context.Context, model.Paging, []model.InternalID,
		[]libauth.UserType, []UserStatus, []model.InternalID, *model.InternalID) ([]*User, int64, error)
	CreateAccount(context.Context, Account, model.InternalID) error
	UpdateAccount(context.Context, Account) error
	UnLinkAccount(context.Context, Account, model.InternalID) error
	ListLinkAccount(context.Context, model.Paging, model.InternalID) ([]*Account, int64, error)
}

type Tiphereth struct {
	auth        *libauth.Auth
	repo        TipherethRepo
	mapper      mapper.LibrarianMapperServiceClient
	searcher    searcher.LibrarianSearcherServiceClient
	porter      porter.LibrarianPorterServiceClient
	pullAccount *libmq.TopicImpl[PullAccountInfo]
}

func NewTiphereth(
	repo TipherethRepo,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	pullAccount *libmq.TopicImpl[PullAccountInfo],
) (*Tiphereth, error) {
	return &Tiphereth{
		auth:        auth,
		repo:        repo,
		mapper:      mClient,
		porter:      pClient,
		searcher:    sClient,
		pullAccount: pullAccount,
	}, nil
}

func (t *Tiphereth) GetToken(ctx context.Context, user *User) (AccessToken, RefreshToken, *errors.Error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.PassWord = password

	user, err = t.repo.FetchUserByPassword(ctx, user)
	if err != nil {
		logger.Infof("FetchUserByPassword failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	if user.Status != UserStatusActive {
		return "", "", pb.ErrorErrorReasonBadRequest("user not active")
	}

	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(user.ID,
		libauth.ClaimsTypeAccessToken, user.Type, nil, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(user.ID,
		libauth.ClaimsTypeRefreshToken, user.Type, nil, time.Hour*24*7) //nolint:gomnd //TODO
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	// TODO save refreshToken to sql
	return AccessToken(accessToken), RefreshToken(refreshToken), nil
}

func (t *Tiphereth) RefreshToken(ctx context.Context) (AccessToken, RefreshToken, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal, libauth.UserTypeSentinel) {
		return "", "", pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return "", "", pb.ErrorErrorReasonUnauthorized("empty token")
	}
	var accessToken, refreshToken string
	accessToken, err := t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeAccessToken, claims.UserType, nil, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeRefreshToken, claims.UserType, nil, time.Hour*24*7) //nolint:gomnd //TODO
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return AccessToken(accessToken), RefreshToken(refreshToken), nil
}

func (t *Tiphereth) CreateDefaultAdmin(ctx context.Context, user *User) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return
	}
	user.PassWord = password
	resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return
	}
	user.ID = model.InternalID(resp.Id)
	user.Status = UserStatusActive
	user.Type = libauth.UserTypeAdmin
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(user.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return
	}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return
	}
}

func (t *Tiphereth) CreateUser(ctx context.Context, user *User) (*model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		if user.Type != libauth.UserTypeSentinel {
			return nil, pb.ErrorErrorReasonForbidden("no permission")
		}
	}
	var creator model.InternalID
	if c, ok := libauth.FromContext(ctx); ok {
		creator = c.InternalID
	}
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.PassWord = password
	resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	user.ID = model.InternalID(resp.Id)
	user.Status = UserStatusActive
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(user.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = t.repo.CreateUser(ctx, user, creator); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := user.ID
	return &res, nil
}

func (t *Tiphereth) UpdateUser( //nolint:gocognit // TODO
	ctx context.Context, user *User, originPassword string,
) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, ok := libauth.FromContext(ctx)
	if !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if user.ID == 0 {
		return pb.ErrorErrorReasonBadRequest("internal id required")
	}
	if user.PassWord != "" && originPassword == "" {
		return pb.ErrorErrorReasonBadRequest("password required")
	}
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) &&
		claims.InternalID != user.ID {
		res, _, err := t.repo.ListUser(ctx,
			model.Paging{
				PageSize: 1,
				PageNum:  1,
			},
			[]model.InternalID{user.ID},
			[]libauth.UserType{libauth.UserTypeSentinel},
			nil,
			nil,
			&claims.InternalID,
		)
		if err != nil || len(res) != 1 || res[0].ID != user.ID {
			return pb.ErrorErrorReasonForbidden("no permission")
		}
	}
	if user.PassWord != "" {
		password, err := t.auth.GeneratePassword(user.PassWord)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
		user.PassWord = password
		originPassword, err = t.auth.GeneratePassword(originPassword)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
	}
	err := t.repo.UpdateUser(ctx, user, originPassword)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListUser(
	ctx context.Context, paging model.Paging, types []libauth.UserType, statuses []UserStatus,
) ([]*User, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	var exclude []model.InternalID
	var creator *model.InternalID
	if c, ok := libauth.FromContext(ctx); !ok {
		return nil, 0, pb.ErrorErrorReasonBadRequest("token required")
	} else {
		if c.UserType != libauth.UserTypeAdmin {
			creator = &c.InternalID
		}
		exclude = append(exclude, c.InternalID)
	}
	users, total, err := t.repo.ListUser(ctx, paging, nil, types, statuses, exclude, creator)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return users, total, nil
}

func (t *Tiphereth) LinkAccount(ctx context.Context, a Account) (*Account, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("invalid token")
	}
	if resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{}); err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	} else {
		a.ID = model.InternalID(resp.Id)
	}
	if _, err := t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(a.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ENTITY,
			Prop: nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if _, err := t.mapper.InsertEdge(ctx, &mapper.InsertEdgeRequest{EdgeList: []*mapper.Edge{
		{
			SrcVid: int64(claims.InternalID),
			DstVid: int64(a.ID),
			Type:   mapper.EdgeType_EDGE_TYPE_EQUAL,
			Prop:   nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err := t.repo.CreateAccount(ctx, a, claims.InternalID); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err := t.pullAccount.Publish(ctx, PullAccountInfo{
		ID:                a.ID,
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
	}); err != nil {
		logger.Errorf("Publish PullAccountInfo failed %s", err.Error())
	}
	return &a, nil
}

func (t *Tiphereth) UnLinkAccount(ctx context.Context, a Account) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("invalid token")
	}
	if err := t.repo.UnLinkAccount(ctx, a, claims.InternalID); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListLinkAccount(
	ctx context.Context, paging model.Paging, id model.InternalID,
) ([]*Account, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("invalid token")
	}
	if id == 0 {
		id = claims.InternalID
	}
	a, total, err := t.repo.ListLinkAccount(ctx, paging, id)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, total, nil
}
