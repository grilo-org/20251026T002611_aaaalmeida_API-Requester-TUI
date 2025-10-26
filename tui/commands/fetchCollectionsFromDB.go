package commands

import (
	"api-requester/context"
	"api-requester/domain/collection"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func FetchCollectionsFromDBCmd(ctx *context.AppContext) tea.Cmd {
	return func() tea.Msg {
		cols, err := collection.GetAllCollection(ctx)
		return msg.LoadCollectionsMsg{
			Collections: cols,
			Err:         err,
		}
	}
}
