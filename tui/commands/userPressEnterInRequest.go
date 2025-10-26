package commands

import (
	"api-requester/domain/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UserPressEnterInRequestCmd(request *request.Request) tea.Cmd {
	return func() tea.Msg {
		return msg.LoadRequestMsg{
			Request: request,
			Err:     nil,
		}
	}
}
