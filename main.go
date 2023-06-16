package main

import (
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/guilehm/solid-robot/internal/services"
	"github.com/guilehm/solid-robot/internal/storages"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		queries.Module,
		services.Module,
		storages.Module,
	).Run()
}
