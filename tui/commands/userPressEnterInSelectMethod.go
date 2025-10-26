package commands

import (
	"api-requester/shared/selectable"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UserPressEnterInSelectCmd(obj selectable.Selectable) tea.Cmd {
	return func() tea.Msg {
		return msg.SendSelectMsg{
			Label: obj.Label(),
			Value: obj.Value(),
		}
	}
}
