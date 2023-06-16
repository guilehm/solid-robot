package main

import (
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/guilehm/solid-robot/internal/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		queries.Module,
		services.Module,
	).Run()
}
