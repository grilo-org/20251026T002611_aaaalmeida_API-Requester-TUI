package button

type Model struct {
	text      string
	isFocused bool
}

func NewModel(text string) *Model {
	return &Model{
		text:      text,
		isFocused: false,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
