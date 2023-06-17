package storages

import (
	postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"
	"go.uber.org/fx"
)

var Module = fx.Options(

	postgresStorage.Module,
	fx.Provide(
		newStorageGroup,
	),
)
