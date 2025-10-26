package request_response_box

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	text := m.body
	if text == nil {
		aux := "body vazio"
		text = &aux
	}
	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(65).
		Height(25).
		Render(*text)

	return bodyBox
}
