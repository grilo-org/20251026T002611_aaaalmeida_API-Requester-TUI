package commands

import (
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func ComponentLoseFocusCmd() tea.Cmd {
	return func() tea.Msg {
		return msg.LoseFocusMsg{}
	}
}
