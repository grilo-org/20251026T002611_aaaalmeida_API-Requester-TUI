package input

import (
	cmds "api-requester/tui/commands"
	messages "api-requester/tui/messages"
	"api-requester/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// USER PRESS ENTER IN HEADER_COMPONENT OR USER PRESS ENTER IN COLLECTION_MENU
	// Receive url and set in input
	case messages.SendStringMsg:
		m.textInput.SetValue(msg.Value)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, cmds.InputChangedCmd(m.textInput.Value())

		case "backspace":
			m.textInput.SetValue(utils.RemoveLastChar(m.textInput.Value()))

		default:
			if utils.IsValidUrlChar(msg.String()) {
				m.textInput.SetValue(m.textInput.Value() + msg.String())
				return m, cmds.InputChangedCmd(m.textInput.Value())
			}
		}
	}

	// component lost focus, so we save changes
	if !m.isFocused {
		return m, cmds.InputChangedCmd(m.textInput.Value())
	}

	return m, nil
}
