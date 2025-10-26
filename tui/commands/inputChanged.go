package commands

import (
	msgs "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func InputChangedCmd(value string) tea.Cmd {
	return func() tea.Msg {
		return msgs.InputChangedMsg{
			Value: value,
		}
	}
}
