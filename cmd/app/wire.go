//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"app/configs"
	"app/internal/infra"
	"app/internal/service"
	"app/internal/usecase"
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
			usecase.ProviderSet,
			service.ProviderSet,
			newApp,
			wire.Bind(new(usecase.GreeterRepo), new(*infra.GreeterRepo)),
		),
	)
}
