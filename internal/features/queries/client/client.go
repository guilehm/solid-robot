package clientQuery

import postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"

type ClientQuery struct {
	postgres *postgresStorage.Postgres
}

func NewClientQuery(postgres *postgresStorage.Postgres) *ClientQuery {
	return &ClientQuery{postgres: postgres}
}
