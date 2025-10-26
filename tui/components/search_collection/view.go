package search_collection

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(30).
		Height(3).
		Align(lipgloss.Left, lipgloss.Center).
		Render(m.textInput.View())
}
