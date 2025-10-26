package main_page

import (
	cmd "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds,
		cmd.FetchCollectionsFromDBCmd(m.context),
		m.subcomponents[REQUEST_HEADERS_INDEX].Init())

	return tea.Batch(cmds...)
}
