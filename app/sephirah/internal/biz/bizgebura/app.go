package bizgebura

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(ctx context.Context, app *modelgebura.App) (*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	app.ID = converter.ToBizInternalID(resp.Id)
	app.Source = modelgebura.AppSourceInternal
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
	return app, nil
}

func (g *Gebura) UpdateApp(ctx context.Context, app *modelgebura.App) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	app.Source = modelgebura.AppSourceInternal
	err := g.repo.UpdateApp(ctx, app)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListApps(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppSource,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
) ([]*modelgebura.App, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, total, err := g.repo.ListApps(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, total, nil
}

func (g *Gebura) MergeApps(ctx context.Context, base modelgebura.App, merged model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if base.Source != modelgebura.AppSourceInternal {
		return pb.ErrorErrorReasonBadRequest("source must be INTERNAL")
	}
	if err := g.repo.MergeApps(ctx, base, merged); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	return nil
}

func (g *Gebura) SearchApps(ctx context.Context, paging model.Paging, keyword string) (
	[]*modelgebura.App, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, total, err := g.repo.SearchApps(ctx, paging, keyword)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	return apps, total, nil
}

func (g *Gebura) GetApp(ctx context.Context, id model.InternalID) (*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, err := g.repo.GetBindApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	var internalApp *modelgebura.App
	var steamApp *modelgebura.App
	for _, app := range apps {
		if app.Source == modelgebura.AppSourceInternal {
			internalApp = app
		}
		if app.Source == modelgebura.AppSourceSteam {
			steamApp = app
		}
	}
	if len(internalApp.Name) == 0 {
		internalApp.Name = steamApp.Name
	}
	if len(internalApp.ShortDescription) == 0 {
		internalApp.ShortDescription = steamApp.ShortDescription
	}
	if len(internalApp.IconImageURL) == 0 {
		internalApp.IconImageURL = steamApp.IconImageURL
	}
	if len(internalApp.HeroImageURL) == 0 {
		internalApp.HeroImageURL = steamApp.HeroImageURL
	}
	if internalApp.Details == nil { //nolint:nestif //TODO
		internalApp.Details = steamApp.Details
	} else if steamApp.Details != nil {
		if len(internalApp.Details.Description) == 0 {
			internalApp.Details.Description = steamApp.Details.Description
		}
		if len(internalApp.Details.ReleaseDate) == 0 {
			internalApp.Details.ReleaseDate = steamApp.Details.ReleaseDate
		}
		if len(internalApp.Details.Developer) == 0 {
			internalApp.Details.Developer = steamApp.Details.Developer
		}
		if len(internalApp.Details.Publisher) == 0 {
			internalApp.Details.Publisher = steamApp.Details.Publisher
		}
		if len(internalApp.Details.Version) == 0 {
			internalApp.Details.Version = steamApp.Details.Version
		}
	}
	return internalApp, nil
}

func (g *Gebura) GetBindApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, err := g.repo.GetBindApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *Gebura) PurchaseApp(ctx context.Context, id model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.PurchaseApp(ctx, claims.InternalID, id)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) GetPurchasedApps(ctx context.Context) ([]*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	} else {
		apps, err := g.repo.GetPurchasedApps(ctx, claims.InternalID)
		if err != nil {
			return nil, pb.ErrorErrorReasonUnspecified("%s", err)
		}
		return apps, nil
	}
}
