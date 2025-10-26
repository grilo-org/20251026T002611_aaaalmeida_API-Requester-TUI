package commands

import (
	"api-requester/context"
	"api-requester/domain/request"

	tea "github.com/charmbracelet/bubbletea"
)

func UpdateRequestCmd(ctx *context.AppContext, req *request.Request) tea.Cmd {
	err := request.UpdateRequest(ctx, req.ID, req)
	if err != nil {
		ctx.Logger.Println(err)
	}
	return nil
}
