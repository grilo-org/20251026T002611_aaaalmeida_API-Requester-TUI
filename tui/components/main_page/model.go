package main_page

import (
	"api-requester/context"
	"api-requester/domain/collection"
	"api-requester/tui/components/collection_menu"
	"api-requester/tui/components/header"
	"api-requester/tui/components/request_header_box"
	"api-requester/tui/components/request_response_box"
	"api-requester/tui/components/search_collection"

	tea "github.com/charmbracelet/bubbletea"
)

// components index
const (
	SEARCH_COLLECTION_INDEX = iota
	COLLECTION_MENU_INDEX
	HEADER_INDEX
	REQUEST_HEADERS_INDEX
	REQUEST_RESPONSE_INDEX
)

// Main TUI component.
// encapsulate everything
type model struct {
	active_component_index int
	subcomponents          []tea.Model
	collections            []*collection.Collection
	context                *context.AppContext
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:                ctx,
		collections:            nil,
		active_component_index: COLLECTION_MENU_INDEX,
		subcomponents: []tea.Model{
			search_collection.NewModel(ctx),
			collection_menu.NewModel(ctx),
			header.NewModel(ctx),
			request_header_box.NewModel(ctx),
			request_response_box.NewModel(ctx),
		},
	}
}
