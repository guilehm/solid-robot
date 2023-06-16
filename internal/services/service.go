package services

import "github.com/guilehm/solid-robot/internal/features/queries"

type ServiceGroup struct {
	queries *queries.QueryGroup
}

func newServiceGroup(queries *queries.QueryGroup) *ServiceGroup {
	return &ServiceGroup{
		queries: queries,
	}
}
