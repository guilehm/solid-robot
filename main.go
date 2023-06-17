package main

import (
	"github.com/guilehm/solid-robot/internal/features/commands"
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/guilehm/solid-robot/internal/logger"
	"github.com/guilehm/solid-robot/internal/services"
	"github.com/guilehm/solid-robot/internal/storages"
	postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"
	"go.uber.org/fx"
)

func main() {

	fx.New(

		// adapters
		logger.Module,

		// internal modules
		queries.Module,
		commands.Module,
		services.Module,
		postgresStorage.Module,

		// listeners
		storages.Module,
		fx.Invoke(
			services.StartListener,
		),
	).Run()
}
