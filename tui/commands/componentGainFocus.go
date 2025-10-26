package commands

import (
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func ComponentGainFocusCmd() tea.Cmd {
	return func() tea.Msg {
		return msg.GainFocusMsg{}
	}
}
