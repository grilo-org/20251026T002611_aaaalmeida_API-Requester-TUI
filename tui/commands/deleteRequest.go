package commands

import (
	"api-requester/context"
	"api-requester/domain/request"
	msg "api-requester/tui/messages"
	"api-requester/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func DeleteRequestCmd(ctx *context.AppContext, requests *[]*request.Request, i int) tea.Cmd {
	req := (*requests)[i]
	err := request.DeleteRequestById(ctx, req.ID)
	if err == nil {
		*requests = utils.RemoveFromArray(*requests, i)
	}
	return func() tea.Msg {
		return msg.ErrorMsg{
			Err: err,
		}
	}
}
