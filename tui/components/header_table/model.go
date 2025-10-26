package header_table

import (
	"api-requester/context"

	"github.com/charmbracelet/bubbles/textinput"
)

type HeaderRow struct {
	KeyTextInput   textinput.Model
	ValueTextInput textinput.Model
}

type Model struct {
	rows        []HeaderRow
	cursorIndex int
	keyOrValue  bool // true = key, false = value
	isFocused   bool
	context     *context.AppContext
}

func NewModel(ctx *context.AppContext) *Model {
	k := textinput.New()
	v := textinput.New()

	return &Model{
		rows: []HeaderRow{
			{KeyTextInput: k, ValueTextInput: v},
		},
		isFocused:   false,
		cursorIndex: 0,
		keyOrValue:  true,
		context:     ctx,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
