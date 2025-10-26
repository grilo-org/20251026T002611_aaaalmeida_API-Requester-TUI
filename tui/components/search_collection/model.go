package search_collection

import (
	"api-requester/context"

	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	context   *context.AppContext
	textInput textinput.Model
	isFocused bool
}

func NewModel(ctx *context.AppContext) model {
	ti := textinput.New()
	ti.Placeholder = "🔍"
	ti.Focus()

	return model{
		context:   ctx,
		textInput: ti,
		isFocused: false,
	}
}
