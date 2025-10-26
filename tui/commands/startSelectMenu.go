package commands

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func StartSelectMenuCmd(s string) tea.Cmd {
	return func() tea.Msg {
		return messages.SendStringMsg{
			Err:   nil,
			Value: s,
		}
	}
}
