package request_response_box

import "api-requester/context"

type Model struct {
	// request *request.Request
	body *string
	ctx  *context.AppContext
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		body: nil,
		ctx:  ctx,
	}
}
