package services

import "github.com/guilehm/solid-robot/internal/features/queries"

type ServiceGroup struct {
	queries       *queries.QueryGroup
	rawMsgChannel chan []string
	bulkAmount    int
}

func newServiceGroup(queries *queries.QueryGroup) *ServiceGroup {
	rawMsgChannel := make(chan []string)
	const bulkAmount = 1000

	return &ServiceGroup{
		queries:       queries,
		rawMsgChannel: rawMsgChannel,
		bulkAmount:    bulkAmount,
	}
}
