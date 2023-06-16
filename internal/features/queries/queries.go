package queries

import (
	clientQuery "github.com/guilehm/solid-robot/internal/features/queries/client"
	postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"
)

type QueryGroup struct {
	Client   *clientQuery.ClientQuery
	Postgres *postgresStorage.Postgres
}

func newQueryGroup(postgres *postgresStorage.Postgres) *QueryGroup {
	return &QueryGroup{
		Client:   clientQuery.NewClientQuery(postgres),
		Postgres: postgres,
	}
}
