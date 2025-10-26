package select_menu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	normalStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(8).Align(lipgloss.Center)
	selectedStyle := lipgloss.NewStyle().Bold(true).BorderForeground(lipgloss.Color("2")).Foreground(lipgloss.Color("2"))

	// empty
	if len(m.Options) == 0 {
		if m.isFocused {
			return normalStyle.Render("---")
		} else {
			return normalStyle.Inherit(selectedStyle).Render("---")
		}
	}

	if !m.isOpened {
		if m.isFocused {
			return normalStyle.Inherit(selectedStyle).Render(m.Options[m.selectedItem].Label())
		} else {
			return normalStyle.Render(m.Options[m.selectedItem].Label())
		}
	}

	var b strings.Builder
	for i, opt := range m.Options {
		if i == m.cursor {
			b.WriteString(selectedStyle.Render(opt.Label()))
		} else {
			b.WriteString(opt.Label())
		}

		if i != len(m.Options)-1 {
			b.WriteString("\n")
		}
	}

	return normalStyle.Render(b.String())
}
