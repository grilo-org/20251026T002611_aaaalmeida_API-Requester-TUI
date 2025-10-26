package button

import (
	cmd "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
		return m, cmd.ButtonPressedCmd(m.text)
	}

	return m, nil
}
