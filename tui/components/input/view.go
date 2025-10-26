package input

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	if m.isFocused {
		return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("2")).
			Foreground(lipgloss.Color("2")).
			Render(m.textInput.View())
	} else {
		return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(m.textInput.View())
	}
}
