package commands

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func SendStringToInputCmd(s string) tea.Cmd {
	return func() tea.Msg {
		return messages.SendStringMsg{
			Err:   nil,
			Value: s,
		}
	}
}
