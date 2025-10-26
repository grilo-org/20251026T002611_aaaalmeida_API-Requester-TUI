package commands

import (
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func ButtonPressedCmd(action string) tea.Cmd {
	return func() tea.Msg {
		return msg.ButtonPressedMsg{
			Action: action,
		}
	}
}
