//go:build wireinject

package main

import (
	grpc2 "github.com/dadaxiaoxiao/account/internal/grpc"
	"github.com/dadaxiaoxiao/account/internal/repository"
	"github.com/dadaxiaoxiao/account/internal/repository/dao"
	"github.com/dadaxiaoxiao/account/internal/service"
	"github.com/dadaxiaoxiao/account/ioc"
	"github.com/dadaxiaoxiao/go-pkg/customserver"
	"github.com/google/wire"
)

var thirdPartyProvider = wire.NewSet(
	ioc.InitDB,
	ioc.InitRedis,
	ioc.InitEtcdClient,
	ioc.InitLogger,
)

func InitApp() *customserver.App {
	wire.Build(
		thirdPartyProvider,
		dao.NewAccountGORMDAO,
		repository.NewAccountRepository,
		service.NewAccountService,
		grpc2.NewAccountServiceServer,
		ioc.InitGRPCServer,
		wire.Struct(new(customserver.App), "GRPCServer"))
	return new(customserver.App)
}
