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
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/client"
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
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, settings *libapp.Settings) (*kratos.App, func(), error) {
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
	tipherethRepo := data.NewTipherethRepo(dataData)
	geburaRepo := data.NewGeburaRepo(dataData)
	yesodRepo := data.NewYesodRepo(dataData)
	librarianMapperServiceClient, err := client.NewMapperClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceClient, err := client.NewPorterClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceClient, err := client.NewSearcherClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	angelaBase, err := bizangela.NewAngelaBase(tipherethRepo, geburaRepo, yesodRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	topic := bizangela.NewPullSteamAppTopic(angelaBase)
	libmqTopic := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase, topic)
	topic2 := bizangela.NewPullAccountTopic(angelaBase, libmqTopic)
	netzachRepo := data.NewNetzachRepo(dataData)
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	libcacheMap := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map2 := bizangela.NewFeedToNotifyFlowMap(netzachRepo, store)
	map3 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic3 := bizangela.NewNotifyPushTopic(angelaBase, map3)
	topic4 := bizangela.NewNotifyRouterTopic(angelaBase, libcacheMap, map2, topic3)
	topic5 := bizangela.NewPullFeedTopic(angelaBase, topic4)
	angela, err := bizangela.NewAngela(libmqMQ, topic2, libmqTopic, topic, topic5, topic4, topic3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	controlBlock := bizbinah.NewControlBlock(libauthAuth)
	binah := bizbinah.NewBinah(controlBlock, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	cron := libcron.NewCron()
	yesod, err := bizyesod.NewYesod(yesodRepo, cron, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic5)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzach := biznetzach.NewNetzach(netzachRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, map2, libcacheMap, map3)
	chesedRepo := data.NewChesedRepo(dataData)
	map4 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, controlBlock, map4)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(libauthAuth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, settings, libauthAuth, v)
	grpcServer := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer, settings)
	httpServer := server.NewGrpcWebServer(grpcServer, sephirah_Server, libauthAuth, settings)
	app := newApp(grpcServer, httpServer, libmqMQ, cron)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
