package storages

import (
	postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"
)

type StorageGroup struct {
	Postgres *postgresStorage.Postgres
}

func newStorageGroup() *StorageGroup {
	return &StorageGroup{
		Postgres: postgresStorage.NewPostgresStorage(),
	}
}
