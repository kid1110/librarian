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
	librarianPorterServiceClient, err := client.NewPorterClient(consul, porter, settings)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	clientPorter, err := client.NewPorter(librarianPorterServiceClient, consul, porter)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	netzachRepo := data.NewNetzachRepo(dataData)
	searcher := client.NewSearcher(librarianSearcherServiceClient)
	topic := biznetzach.NewSystemNotificationTopic(netzachRepo, searcher)
	tipherethRepo := data.NewTipherethRepo(dataData)
	libcacheMap := biztiphereth.NewPorterInstanceCache(tipherethRepo, store)
	map2 := biztiphereth.NewPorterContextCache(tipherethRepo, store)
	supervisorSupervisor, err := supervisor.NewSupervisor(porter, mq, auth, clientPorter, topic, libcacheMap, map2)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, supervisorSupervisor, geburaRepo, librarianPorterServiceClient, searcher)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	map3 := bizangela.NewAppInfoCache(geburaRepo, store)
	libmqTopic := bizangela.NewUpdateAppInfoIndexTopic(angelaBase)
	topic2 := bizangela.NewPullAppInfoTopic(angelaBase, map3, libmqTopic)
	topic3 := bizangela.NewPullAccountAppInfoRelationTopic(angelaBase, topic2)
	topic4 := bizangela.NewPullAccountTopic(angelaBase, topic3)
	map4 := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map5 := bizangela.NewFeedToNotifyFlowCache(netzachRepo, store)
	map6 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic5 := bizangela.NewNotifyPushTopic(angelaBase, map6)
	topic6 := bizangela.NewNotifyRouterTopic(angelaBase, map4, map5, topic5)
	topic7 := bizangela.NewFeedItemPostprocessTopic(angelaBase, topic6, topic)
	topic8 := bizangela.NewPullFeedTopic(angelaBase, topic7, topic)
	angela, err := bizangela.NewAngela(mq, topic4, topic3, topic2, topic8, topic6, topic5, topic7, libmqTopic)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	key := biztiphereth.NewUserCountCache(tipherethRepo, store)
	tiphereth, err := biztiphereth.NewTiphereth(settings, tipherethRepo, auth, supervisorSupervisor, searcher, topic4, cron, key, libcacheMap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, auth, searcher, librarianPorterServiceClient, supervisorSupervisor, libmqTopic, topic2, map3)
	binahRepo, err := data.NewBinahRepo(s3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	controlBlock := bizbinah.NewControlBlock(auth)
	binah := bizbinah.NewBinah(binahRepo, controlBlock, auth, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	map7 := bizyesod.NewFeedOwnerCache(yesodRepo, store)
	yesod, err := bizyesod.NewYesod(yesodRepo, supervisorSupervisor, cron, searcher, topic8, topic, map7)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	netzach, err := biznetzach.NewNetzach(netzachRepo, supervisorSupervisor, searcher, mq, map5, map4, map6, topic)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	chesedRepo := data.NewChesedRepo(dataData)
	map8 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, binahRepo, cron, librarianPorterServiceClient, searcher, librarianMinerServiceClient, controlBlock, map8)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, supervisorSupervisor, settings, auth, sephirahServer)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
