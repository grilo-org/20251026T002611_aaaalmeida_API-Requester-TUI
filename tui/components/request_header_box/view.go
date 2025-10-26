package request_header_box

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var topsideContent, bottomsideContent []string

	topsideContent = append(topsideContent,
		m.subcomponents[SELECT_MENU_INDEX].View(),
		m.subcomponents[INPUT_URL_INDEX].View(),
		m.subcomponents[SEND_REQUEST_BUTTON_INDEX].View())
	topsideView := lipgloss.JoinHorizontal(lipgloss.Top, topsideContent...)

	// TODO: adicionar view do body da request depois que estiver pronto
	bottomsideContent = append(bottomsideContent,
		m.subcomponents[HEADER_TABLE_INDEX].View())
	bottomsideView := lipgloss.JoinHorizontal(lipgloss.Top, bottomsideContent...)

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(65).
		Height(25).
		Render(lipgloss.JoinVertical(lipgloss.Top, topsideView, bottomsideView))

	return bodyBox
}
