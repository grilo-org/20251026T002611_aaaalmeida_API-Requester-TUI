package header

import (
	"api-requester/context"
	"api-requester/domain/request"
)

type Model struct {
	context          *context.AppContext
	cursor           int
	selectedTabIndex int
	requests         []*request.Request
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		cursor:           0,
		selectedTabIndex: 0,
		requests:         nil,
		context:          ctx,
	}
}
