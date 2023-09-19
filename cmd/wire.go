//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/google/wire"

	"github/kunhou/simple-backend/deliver/http/controller"
	"github/kunhou/simple-backend/deliver/http/router"
	"github/kunhou/simple-backend/deliver/http/server"
	"github/kunhou/simple-backend/pkg/config"
	"github/kunhou/simple-backend/pkg/data"
	"github/kunhou/simple-backend/pkg/servmanager"
	"github/kunhou/simple-backend/repository"
	"github/kunhou/simple-backend/usecase"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *config.Server,
	dbConf *data.DatabaseConf,
) (*servmanager.Application, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		usecase.ProviderSetUsecase,
		repository.ProviderSetRepository,
		newApplication,
	))
}
