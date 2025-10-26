package header_table

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m *Model) View() string {
	// header style
	headerStyle := lipgloss.NewStyle().Bold(true).Width(25)

	// row style
	baseRowStyle := lipgloss.NewStyle().Width(25)
	oddRowStyle := baseRowStyle.Background(lipgloss.Color("#1E1E1E"))
	evenRowStyle := baseRowStyle.Background(lipgloss.Color("#2A2A2A"))
	selectedRowStyle := baseRowStyle.Background(lipgloss.Color("#3C3C3C"))

	// border style
	var borderStyle lipgloss.Style
	if m.isFocused {
		borderStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("2")).
			Foreground(lipgloss.Color("2"))
	} else {
		borderStyle = lipgloss.NewStyle()
	}

	var rows [][]string
	for i, r := range m.rows {
		row := []string{r.KeyTextInput.Value(), r.ValueTextInput.Value()}

		// each line has a style
		switch {
		case i == m.cursorIndex:
			for j := range row {
				row[j] = selectedRowStyle.Render(row[j])
			}

		case i%2 == 0:
			for j := range row {
				row[j] = evenRowStyle.Render(row[j])
			}

		default:
			for j := range row {
				row[j] = oddRowStyle.Render(row[j])
			}
		}

		rows = append(rows, row)
	}

	return table.New().
		Headers("Key", "Value").
		StyleFunc(
			func(row, col int) lipgloss.Style {
				if row == 0 {
					return headerStyle
				}
				return baseRowStyle
			}).
		Border(lipgloss.NormalBorder()).
		BorderStyle(borderStyle).
		Rows(rows...).
		Render()
}
