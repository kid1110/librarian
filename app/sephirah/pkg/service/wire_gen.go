// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
	v1_2 "github.com/tuihub/protos/pkg/librarian/miner/v1"
	"github.com/tuihub/protos/pkg/librarian/searcher/v1"
	v1_3 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// Injectors from wire.go:

func NewSephirahService(sephirahServer *conf.SephirahServer, database *conf.Database, s3 *conf.S3, porter *conf.Porter, consul *conf.Consul, auth *libauth.Auth, mq *libmq.MQ, cron *libcron.Cron, store libcache.Store, settings *libapp.Settings, librarianSearcherServiceClient v1.LibrarianSearcherServiceClient, librarianMinerServiceClient v1_2.LibrarianMinerServiceClient) (v1_3.LibrarianSephirahServiceServer, func(), error) {
	entClient, cleanup, err := data.NewSQLClient(database, settings)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	angelaRepo := data.NewAngelaRepo(dataData)
	librarianPorterServiceClient, err := client.NewPorterClient(consul)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	clientPorter, err := client.NewPorter(librarianPorterServiceClient, consul)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	supervisorSupervisor, err := supervisor.NewSupervisor(porter, auth, clientPorter)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	searcher := client.NewSearcher(librarianSearcherServiceClient)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, supervisorSupervisor, geburaRepo, librarianPorterServiceClient, searcher)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	libcacheMap := bizangela.NewAppInfoCache(geburaRepo, store)
	topic := bizangela.NewUpdateAppInfoIndexTopic(angelaBase)
	libmqTopic := bizangela.NewPullAppInfoTopic(angelaBase, libcacheMap, topic)
	topic2 := bizangela.NewPullAccountAppInfoRelationTopic(angelaBase, libmqTopic)
	topic3 := bizangela.NewPullAccountTopic(angelaBase, topic2)
	netzachRepo := data.NewNetzachRepo(dataData)
	map2 := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map3 := bizangela.NewFeedToNotifyFlowCache(netzachRepo, store)
	map4 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic4 := bizangela.NewNotifyPushTopic(angelaBase, map4)
	topic5 := bizangela.NewNotifyRouterTopic(angelaBase, map2, map3, topic4)
	topic6 := bizangela.NewParseFeedItemDigestTopic(angelaBase)
	topic7 := bizangela.NewPullFeedTopic(angelaBase, topic5, topic6)
	topic8 := bizangela.NewSystemNotificationTopic(angelaBase)
	angela, err := bizangela.NewAngela(mq, topic3, topic2, libmqTopic, topic7, topic5, topic4, topic6, topic, topic8)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tipherethRepo := data.NewTipherethRepo(dataData)
	key := biztiphereth.NewUserCountCache(tipherethRepo, store)
	tiphereth, err := biztiphereth.NewTiphereth(settings, tipherethRepo, auth, supervisorSupervisor, searcher, topic3, cron, key)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, auth, searcher, librarianPorterServiceClient, supervisorSupervisor, topic, libmqTopic, libcacheMap)
	binahRepo, err := data.NewBinahRepo(s3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	controlBlock := bizbinah.NewControlBlock(auth)
	binah := bizbinah.NewBinah(binahRepo, controlBlock, auth, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	yesod, err := bizyesod.NewYesod(yesodRepo, supervisorSupervisor, cron, searcher, topic7)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	netzach := biznetzach.NewNetzach(netzachRepo, supervisorSupervisor, searcher, map3, map2, map4)
	chesedRepo := data.NewChesedRepo(dataData)
	map5 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, binahRepo, cron, librarianPorterServiceClient, searcher, librarianMinerServiceClient, controlBlock, map5)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(auth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, supervisorSupervisor, settings, auth, v, sephirahServer)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
