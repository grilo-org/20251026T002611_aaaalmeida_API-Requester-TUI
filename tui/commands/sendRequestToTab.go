package commands

import (
	"api-requester/domain/request"
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func SendRequestToTabCmd(req *request.Request) tea.Cmd {
	return func() tea.Msg {
		return messages.SendRequestMsg{
			Request: req,
			Err:     nil,
		}
	}
}
