package queries

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		newQueryGroup,
	),
)
