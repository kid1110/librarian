// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/app/searcher/internal/server"
	"github.com/tuihub/librarian/app/searcher/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(searcher_Server *conf.Searcher_Server, searcher_Data *conf.Searcher_Data, consul *conf.Consul, settings *libapp.Settings) (*kratos.App, func(), error) {
	v, err := data.NewBleve(searcher_Data, settings)
	if err != nil {
		return nil, nil, err
	}
	serviceManager, err := data.NewMeili(searcher_Data, settings)
	if err != nil {
		return nil, nil, err
	}
	sonyflake := data.NewSnowFlake()
	searcherRepo, err := data.NewSearcherRepo(v, serviceManager, sonyflake)
	if err != nil {
		return nil, nil, err
	}
	searcher := biz.NewSearcher(searcherRepo)
	librarianSearcherServiceServer := service.NewLibrarianSearcherServiceService(searcher)
	grpcServer := server.NewGRPCServer(searcher_Server, librarianSearcherServiceServer, settings)
	registrar, err := libapp.NewRegistrar(consul)
	if err != nil {
		return nil, nil, err
	}
	app := newApp(grpcServer, registrar)
	return app, func() {
	}, nil
}
