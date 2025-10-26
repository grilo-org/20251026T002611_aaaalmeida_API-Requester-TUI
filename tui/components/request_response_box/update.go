package request_response_box

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.SendStringMsg:
		m.body = &msg.Value

	case messages.LoadResponseMsg:
		m.body = &msg.Value
	}
	return m, nil
}
