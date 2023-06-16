package queries

import clientQuery "github.com/guilehm/solid-robot/internal/features/queries/client"

type QueryGroup struct {
	Client *clientQuery.ClientQuery
}

func newQueryGroup() *QueryGroup {
	return &QueryGroup{
		Client: clientQuery.NewClientQuery(),
	}
}
