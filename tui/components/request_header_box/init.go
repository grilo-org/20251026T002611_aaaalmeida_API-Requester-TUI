package request_header_box

import (
	command "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return command.FetchMethodsCmd(m.context)
}
