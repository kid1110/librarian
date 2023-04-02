// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
	"github.com/tuihub/protos/pkg/librarian/mapper/v1"
	v1_3 "github.com/tuihub/protos/pkg/librarian/porter/v1"
	v1_2 "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	v1_4 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// Injectors from wire.go:

func NewSephirahService(sephirah_Data *conf.Sephirah_Data, auth *libauth.Auth, mq *libmq.MQ, cron *libcron.Cron, settings *libapp.Settings, librarianMapperServiceClient v1.LibrarianMapperServiceClient, librarianSearcherServiceClient v1_2.LibrarianSearcherServiceClient, librarianPorterServiceClient v1_3.LibrarianPorterServiceClient) (v1_4.LibrarianSephirahServiceServer, func(), error) {
	client, cleanup, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(client)
	tipherethRepo := data.NewTipherethRepo(dataData)
	geburaRepo := data.NewGeburaRepo(dataData)
	yesodRepo := data.NewYesodRepo(dataData)
	angelaBase, err := bizangela.NewAngelaBase(tipherethRepo, geburaRepo, yesodRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	topic := bizangela.NewPullSteamAppTopic(angelaBase)
	libmqTopic := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase, topic)
	topic2 := bizangela.NewPullAccountTopic(angelaBase, libmqTopic)
	topic3 := bizangela.NewPullFeedTopic(angelaBase)
	angela, err := bizangela.NewAngela(mq, topic2, libmqTopic, topic, topic3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic2)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	callbackControlBlock := bizbinah.NewCallbackControl()
	gebura := bizgebura.NewGebura(geburaRepo, auth, callbackControlBlock, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	binah := bizbinah.NewBinah(callbackControlBlock, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	yesod, err := bizyesod.NewYesod(yesodRepo, cron, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topic3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(auth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, settings, v)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
