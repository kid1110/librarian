// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/app/mapper/internal/data"
	"github.com/tuihub/librarian/app/mapper/internal/server"
	"github.com/tuihub/librarian/app/mapper/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(mapper_Server *conf.Mapper_Server, mapper_Data *conf.Mapper_Data, consul *conf.Consul, settings *libapp.Settings) (*kratos.App, func(), error) {
	db, cleanup := data.NewNebula(mapper_Data)
	handle, cleanup2, err := data.NewCayley(mapper_Data, settings)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mapperRepo, err := data.NewMapperRepo(db, handle)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mapper := biz.NewMapper(mapperRepo)
	librarianMapperServiceServer := service.NewLibrarianMapperServiceService(mapper)
	grpcServer := server.NewGRPCServer(mapper_Server, librarianMapperServiceServer, settings)
	registrar, err := libapp.NewRegistrar(consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(grpcServer, registrar)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
