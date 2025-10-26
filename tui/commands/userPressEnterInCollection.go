package commands

import (
	"api-requester/context"
	"api-requester/domain/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UserPressEnterInCollectionCmd(ctx *context.AppContext, collection_id int) tea.Cmd {
	return func() tea.Msg {
		reqs, err := request.SearchRequestByCollectionId(ctx, collection_id)
		return msg.LoadRequestFromCollectionMsg{
			Collection_id: collection_id,
			Requests:      reqs,
			Err:           err,
		}
	}
}
