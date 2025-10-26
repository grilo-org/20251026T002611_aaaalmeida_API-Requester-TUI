package focusable

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	Focus()
	Blur()
	View() string
	Update(tea.Msg) (tea.Model, tea.Cmd)
}
