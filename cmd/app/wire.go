//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"app/configs"
	"app/handler"
	"app/infra"
	"app/infra/database"
	"app/infra/encryption"
	"app/internal/user/port/driven"
	"app/internal/user/port/driver"
	"app/internal/user/usecase"
	"app/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*configs.ApplicationConfig, *configs.DBConfig, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			infra.ProviderSet,
			handler.ProviderSet,
			newApp,
			usecase.NewUserWriterUsecase,
			wire.Bind(new(driven.Encyptor), new(*encryption.BcryptEncryption)),
			wire.Bind(new(driven.UserWriter), new(*database.UserRepository)),
			wire.Bind(new(driver.UserWriterUsecase), new(*usecase.UserWriterUsecase)),
		),
	)
}
