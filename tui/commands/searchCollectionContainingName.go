package commands

import (
	"api-requester/domain/collection"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func SearchCollectionContainingNameCmd(collections []*collection.Collection, collection_name string) tea.Cmd {
	return func() tea.Msg {
		cols := collection.SearchCollectionContainingName(collections, collection_name)

		return msg.LoadCollectionsMsg{
			Collections: cols,
			Err:         nil,
		}
	}
}
