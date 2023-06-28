// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
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
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
	"github.com/tuihub/protos/pkg/librarian/mapper/v1"
	v1_4 "github.com/tuihub/protos/pkg/librarian/miner/v1"
	v1_3 "github.com/tuihub/protos/pkg/librarian/porter/v1"
	v1_2 "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	v1_5 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// Injectors from wire.go:

func NewSephirahService(sephirah_Data *conf.Sephirah_Data, auth *libauth.Auth, mq *libmq.MQ, cron *libcron.Cron, store libcache.Store, settings *libapp.Settings, librarianMapperServiceClient v1.LibrarianMapperServiceClient, librarianSearcherServiceClient v1_2.LibrarianSearcherServiceClient, librarianPorterServiceClient v1_3.LibrarianPorterServiceClient, librarianMinerServiceClient v1_4.LibrarianMinerServiceClient) (v1_5.LibrarianSephirahServiceServer, func(), error) {
	client, cleanup, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(client)
	angelaRepo := data.NewAngelaRepo(dataData)
	geburaRepo := data.NewGeburaRepo(dataData)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, geburaRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	topic := bizangela.NewPullSteamAppTopic(angelaBase)
	libmqTopic := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase, topic)
	topic2 := bizangela.NewPullAccountTopic(angelaBase, libmqTopic)
	netzachRepo := data.NewNetzachRepo(dataData)
	libcacheMap := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map2 := bizangela.NewFeedToNotifyFlowMap(netzachRepo, store)
	map3 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic3 := bizangela.NewNotifyPushTopic(angelaBase, map3)
	topic4 := bizangela.NewNotifyRouterTopic(angelaBase, libcacheMap, map2, topic3)
	topic5 := bizangela.NewPullFeedTopic(angelaBase, topic4)
	angela, err := bizangela.NewAngela(mq, topic2, libmqTopic, topic, topic5, topic4, topic3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tipherethRepo := data.NewTipherethRepo(dataData)
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic2)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	controlBlock := bizbinah.NewControlBlock(auth)
	binah := bizbinah.NewBinah(controlBlock, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	yesod, err := bizyesod.NewYesod(yesodRepo, cron, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic5)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	netzach := biznetzach.NewNetzach(netzachRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, map2, libcacheMap, map3)
	chesedRepo := data.NewChesedRepo(dataData)
	map4 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, cron, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, librarianMinerServiceClient, controlBlock, map4)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(auth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, settings, auth, v)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
