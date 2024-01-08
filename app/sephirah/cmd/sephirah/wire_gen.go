// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
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
	client2 "github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, sephirah_Porter *conf.Sephirah_Porter, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, settings *libapp.Settings) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	libmqMQ, cleanup, err := libmq.NewMQ(mq, settings)
	if err != nil {
		return nil, nil, err
	}
	entClient, cleanup2, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	angelaRepo := data.NewAngelaRepo(dataData)
	librarianPorterServiceClient, err := client.NewPorterClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	porter, err := client.NewPorter(librarianPorterServiceClient)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cron := libcron.NewCron()
	supervisorSupervisor, err := supervisor.NewSupervisor(sephirah_Porter, porter, cron)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	librarianMapperServiceClient, err := client2.NewMapperClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceClient, err := client2.NewSearcherClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	searcher := client.NewSearcher(librarianSearcherServiceClient)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, supervisorSupervisor, geburaRepo, librarianMapperServiceClient, librarianPorterServiceClient, searcher)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	topic := bizangela.NewUpdateAppIndexTopic(angelaBase)
	libmqTopic := bizangela.NewPullSteamAppTopic(angelaBase, topic)
	topic2 := bizangela.NewPullAccountAppRelationTopic(angelaBase, libmqTopic)
	topic3 := bizangela.NewPullAccountTopic(angelaBase, topic2)
	netzachRepo := data.NewNetzachRepo(dataData)
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	libcacheMap := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map2 := bizangela.NewFeedToNotifyFlowCache(netzachRepo, store)
	map3 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic4 := bizangela.NewNotifyPushTopic(angelaBase, map3)
	topic5 := bizangela.NewNotifyRouterTopic(angelaBase, libcacheMap, map2, topic4)
	topic6 := bizangela.NewParseFeedItemDigestTopic(angelaBase)
	topic7 := bizangela.NewPullFeedTopic(angelaBase, topic5, topic6)
	angela, err := bizangela.NewAngela(libmqMQ, topic3, topic2, libmqTopic, topic7, topic5, topic4, topic6, topic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tipherethRepo := data.NewTipherethRepo(dataData)
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, libauthAuth, supervisorSupervisor, librarianMapperServiceClient, searcher, topic3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, librarianMapperServiceClient, searcher, topic)
	binahRepo, err := data.NewBinahRepo(sephirah_Data)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	controlBlock := bizbinah.NewControlBlock(libauthAuth)
	binah := bizbinah.NewBinah(binahRepo, controlBlock, libauthAuth, librarianMapperServiceClient, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	yesod, err := bizyesod.NewYesod(yesodRepo, supervisorSupervisor, cron, librarianMapperServiceClient, searcher, topic7)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzach := biznetzach.NewNetzach(netzachRepo, supervisorSupervisor, searcher, map2, libcacheMap, map3)
	chesedRepo := data.NewChesedRepo(dataData)
	librarianMinerServiceClient, err := client2.NewMinerClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	map4 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, binahRepo, cron, librarianPorterServiceClient, searcher, librarianMinerServiceClient, controlBlock, map4)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(libauthAuth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, supervisorSupervisor, settings, libauthAuth, v)
	grpcServer, err := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer := server.NewGrpcWebServer(grpcServer, sephirah_Server, libauthAuth, settings)
	app := newApp(grpcServer, httpServer, libmqMQ, cron)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
