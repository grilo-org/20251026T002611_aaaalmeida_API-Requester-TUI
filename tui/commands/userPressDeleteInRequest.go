package commands

import (
	"api-requester/domain/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UserPressDeleteInRequestCmd(req *request.Request) tea.Cmd {
	return func() tea.Msg {
		return msg.DeleteRequestMsg{
			Request: req,
		}
	}
}
