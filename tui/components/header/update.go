package header

import (
	commands "api-requester/tui/commands"
	messages "api-requester/tui/messages"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "delete", "backspace":
			return m, commands.UserPressDeleteInRequestCmd(m.requests[m.cursor])
		case "enter", " ":
			m.selectedTabIndex = m.cursor
			return m, commands.UserPressEnterInRequestCmd(m.requests[m.selectedTabIndex])

		case "right":
			if m.cursor == len(m.requests)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}
		case "left":
			if m.cursor == 0 {
				m.cursor = len(m.requests) - 1
			} else {
				m.cursor--
			}

			// case "delete", "backspace":
			// m.requests = append(m.requests, )
		}

		// USER PRESS ENTER IN COLLECTION MENU
		// ADD REQUEST ONLY IF NOT ALREADY EXISTS
	case messages.SendRequestMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.requests = append(m.requests, msg.Request)
		}
	}

	return m, nil
}
