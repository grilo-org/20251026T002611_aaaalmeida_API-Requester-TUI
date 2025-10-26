package input

import (
	"api-requester/context"

	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput textinput.Model
	isFocused bool
	context   *context.AppContext
}

func NewModel(width int, placeholder *string, ctx *context.AppContext) *Model {
	ti := textinput.New()
	ti.Prompt = "# "
	ti.Width = width
	if placeholder != nil {
		ti.Placeholder = *placeholder
	}

	return &Model{
		context:   ctx,
		textInput: ti,
		isFocused: false,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
