package bizgebura

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcherpb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(ctx context.Context, app *modelgebura.App) (*modelgebura.App, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return nil, bizutils.NoPermissionError()
	}
	id, err := g.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	app.ID = id
	app.Internal = true
	app.SourceAppID = strconv.FormatInt(int64(app.ID), 10)
	app.SourceURL = ""
	app.BoundInternal = app.ID
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(app.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
		}},
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = g.repo.CreateApp(ctx, app); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{app.ID}})
	return app, nil
}

func (g *Gebura) UpdateApp(ctx context.Context, app *modelgebura.App) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	app.Internal = true
	err := g.repo.UpdateApp(ctx, app)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{app.ID}})
	return nil
}

func (g *Gebura) ListApps(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
) ([]*modelgebura.App, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	apps, total, err := g.repo.ListApps(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, total, nil
}

func (g *Gebura) MergeApps(ctx context.Context, base modelgebura.App, merged model.InternalID) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	if !base.Internal {
		return pb.ErrorErrorReasonBadRequest("source must be INTERNAL")
	}
	if err := g.repo.MergeApps(ctx, base, merged); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{base.ID}})
	return nil
}

func (g *Gebura) SearchApps(ctx context.Context, paging model.Paging, keyword string) (
	[]*modelgebura.AppMixed, int, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	ids, err := g.searcher.SearchID(ctx, paging, keyword, searcherpb.Index_INDEX_GEBURA_APP)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	apps, err := g.repo.GetBatchBoundApps(ctx, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	res := make([]*modelgebura.AppMixed, 0, len(apps))
	for _, a := range apps {
		res = append(res, a.Flatten())
	}
	return res, 0, nil
}

func (g *Gebura) GetApp(ctx context.Context, id model.InternalID) (*modelgebura.App, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	apps, err := g.repo.GetBoundApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	for _, app := range apps {
		if app.ID == id {
			return app, nil
		}
	}
	return nil, pb.ErrorErrorReasonNotFound("app not found")
}

func (g *Gebura) GetBoundApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	apps, err := g.repo.GetBoundApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *Gebura) PurchaseApp(ctx context.Context, id model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	} else {
		err := g.repo.PurchaseApp(ctx, claims.InternalID, id)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) GetPurchasedApps(ctx context.Context) ([]*modelgebura.AppMixed, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	} else {
		apps, err := g.repo.GetPurchasedApps(ctx, claims.InternalID)
		if err != nil {
			return nil, pb.ErrorErrorReasonUnspecified("%s", err)
		}
		res := make([]*modelgebura.AppMixed, 0, len(apps))
		for _, a := range apps {
			res = append(res, a.Flatten())
		}
		return res, nil
	}
}
