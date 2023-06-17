package main

import (
	"github.com/guilehm/solid-robot/internal/features/commands"
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/guilehm/solid-robot/internal/logger"
	"github.com/guilehm/solid-robot/internal/services"
	"github.com/guilehm/solid-robot/internal/storages"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		// internal modules
		logger.Module,
		queries.Module,
		commands.Module,
		services.Module,
		storages.Module,

		// listeners
		fx.Invoke(
			services.StartListener,
		),
	).Run()
}
