package request_header_box

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/shared/focusable"
	"api-requester/tui/components/button"
	"api-requester/tui/components/header_table"
	"api-requester/tui/components/input"
	"api-requester/tui/components/select_menu"
)

const (
	SELECT_MENU_INDEX = iota
	INPUT_URL_INDEX
	SEND_REQUEST_BUTTON_INDEX
	HEADER_TABLE_INDEX
)

type Model struct {
	context                *context.AppContext
	request                *request.Request
	selectedComponentIndex int
	cursor                 int
	subcomponents          []focusable.Focusable
}

func NewModel(ctx *context.AppContext) Model {
	placeholder := "Inform url"
	return Model{
		context:                ctx,
		request:                nil,
		selectedComponentIndex: INPUT_URL_INDEX,
		cursor:                 INPUT_URL_INDEX,
		subcomponents: []focusable.Focusable{
			select_menu.NewModel(ctx),
			input.NewModel(40, &placeholder, ctx),
			button.NewModel(">>"),
			header_table.NewModel(ctx),
		},
	}
}
