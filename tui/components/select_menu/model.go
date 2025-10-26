package select_menu

import (
	"api-requester/context"
	"api-requester/shared/selectable"
)

type Model struct {
	Options      []selectable.Selectable
	cursor       int
	selectedItem int
	isOpened     bool
	isFocused    bool
	ctx          *context.AppContext
}

func NewModel(ctx *context.AppContext) *Model {
	return &Model{
		ctx:          ctx,
		isFocused:    false,
		isOpened:     false,
		cursor:       0,
		selectedItem: 0,
		Options:      nil,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
