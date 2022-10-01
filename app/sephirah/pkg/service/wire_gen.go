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
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/protos/pkg/librarian/mapper/v1"
	v1_3 "github.com/tuihub/protos/pkg/librarian/porter/v1"
	v1_2 "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	v1_4 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// Injectors from wire.go:

func NewSephirahService(sephirah_Data *conf.Sephirah_Data, auth *libauth.Auth, mq *libmq.MQ, librarianMapperServiceClient v1.LibrarianMapperServiceClient, librarianSearcherServiceClient v1_2.LibrarianSearcherServiceClient, librarianPorterServiceClient v1_3.LibrarianPorterServiceClient) (v1_4.LibrarianSephirahServiceServer, func(), error) {
	angelaBase, err := bizangela.NewAngelaBase(librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		return nil, nil, err
	}
	topicImpl := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase)
	libmqTopicImpl := bizangela.NewPullAccountTopic(angelaBase, topicImpl)
	angela, err := bizangela.NewAngela(mq, libmqTopicImpl, topicImpl)
	if err != nil {
		return nil, nil, err
	}
	client, cleanup, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(client)
	tipherethRepo := data.NewTipherethRepo(dataData)
	tipherethUseCase, err := biztiphereth.NewTipherethUseCase(tipherethRepo, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, libmqTopicImpl)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	callbackControlBlock := bizbinah.NewCallbackControl()
	geburaUseCase := bizgebura.NewGeburaUseCase(geburaRepo, auth, callbackControlBlock, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	binahUseCase := bizbinah.NewBinahUseCase(callbackControlBlock, auth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tipherethUseCase, geburaUseCase, binahUseCase)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
