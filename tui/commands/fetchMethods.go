package commands

import (
	"api-requester/context"
	"api-requester/domain/method"
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func FetchMethodsCmd(ctx *context.AppContext) tea.Cmd {
	methods, err := method.GetAllMethod(ctx)
	return func() tea.Msg {
		return messages.LoadMethodsMsg{
			Methods: methods,
			Err:     err,
		}
	}
}
