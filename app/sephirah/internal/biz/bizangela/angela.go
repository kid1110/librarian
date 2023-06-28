package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAngela,
	NewAngelaBase,
	NewPullAccountTopic,
	NewPullSteamAccountAppRelationTopic,
	NewPullSteamAppTopic,
	NewPullFeedTopic,
	NewNotifyRouterTopic,
	NewNotifyPushTopic,
	NewFeedToNotifyFlowMap,
	NewNotifyFlowCache,
	NewNotifyTargetCache,
)

type Angela struct {
	mq *libmq.MQ
}
type AngelaBase struct {
	repo     AngelaRepo
	g        bizgebura.GeburaRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
}

type AngelaRepo interface {
	UpdateAccount(context.Context, modeltiphereth.Account) error
	UpsertApps(context.Context, []*modelgebura.App) error
	UpsertFeed(context.Context, *modelfeed.Feed) error
	UpsertFeedItems(context.Context, []*modelfeed.Item, model.InternalID) ([]string, error)
}

func NewAngelaBase(
	repo AngelaRepo,
	g bizgebura.GeburaRepo,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) (*AngelaBase, error) {
	return &AngelaBase{
		repo:     repo,
		g:        g,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}, nil
}

func NewAngela(
	mq *libmq.MQ,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	pullSteamAccountAppRelation *libmq.Topic[modelangela.PullSteamAccountAppRelation],
	pullSteamApp *libmq.Topic[modelangela.PullSteamApp],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelangela.NotifyRouter],
	notifyPush *libmq.Topic[modelangela.NotifyPush],
) (*Angela, error) {
	if err := mq.RegisterTopic(pullAccount); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullSteamAccountAppRelation); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullSteamApp); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullFeed); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(notifyRouter); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(notifyPush); err != nil {
		return nil, err
	}
	return &Angela{
		mq: mq,
	}, nil
}
