package header

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	selectedStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Underline(true)
	normalStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Faint(true)

	var tabs []string
	for i, req := range m.requests {
		if i == m.cursor {
			tabs = append(tabs, selectedStyle.Render(req.Name))
		} else {
			tabs = append(tabs, normalStyle.Render(req.Name))
		}
	}

	headerContent := lipgloss.JoinHorizontal(lipgloss.Left, tabs...)
	headerBox := lipgloss.NewStyle().Border(lipgloss.ThickBorder()).Width(132).Height(3)
	return headerBox.Render(headerContent)
}
