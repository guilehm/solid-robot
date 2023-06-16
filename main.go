package main

import (
	"github.com/guilehm/solid-robot/internal/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		services.Module,
	).Run()
}
