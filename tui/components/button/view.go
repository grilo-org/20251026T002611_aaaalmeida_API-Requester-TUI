package button

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	if m.isFocused {
		return lipgloss.NewStyle().Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("2")).
			Foreground(lipgloss.Color("2")).
			Render(m.text)
	} else {
		return lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.NormalBorder()).Render(m.text)
	}
}
