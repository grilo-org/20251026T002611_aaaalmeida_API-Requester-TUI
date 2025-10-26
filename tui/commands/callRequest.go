package commands

import (
	"api-requester/domain/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func CallRequestCmd(req *request.Request) tea.Cmd {
	response, err := request.CallRequest(req)
	return func() tea.Msg {
		return msg.LoadResponseMsg{
			Value: response,
			Err:   err,
		}
	}
}
