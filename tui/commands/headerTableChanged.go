package commands

import (
	msgs "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func HeaderTableChangedCmd(content map[string]string) tea.Cmd {
	return func() tea.Msg {
		return msgs.HeaderTableChangedMsg{
			Content: content,
		}
	}
}
