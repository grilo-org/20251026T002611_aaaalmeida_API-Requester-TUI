package collection_menu

import (
	"api-requester/utils"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

// TODO: add style configuration
func (m model) View() string {
	selectedStyle := lipgloss.NewStyle().Bold(true)
	normalStyle := lipgloss.NewStyle().Faint(true)

	t := tree.New()

	for i, col := range m.collections {
		var subTree *tree.Tree
		colName := utils.Truncate(col.Name, 20)

		// cursor in collection
		if m.cursor.colIndex == i && m.cursor.reqIndex == nil {
			subTree = t.Child(selectedStyle.Render(colName))
		} else {
			subTree = t.Child(normalStyle.Render(colName))
		}

		// only loads fetched collections
		if len(col.Requests) > 0 && m.openCloseIndex[i] {
			for j, r := range col.Requests {
				label := utils.Truncate(
					utils.Concatenate(
						utils.TransformMethodIdToVerbColored(r.Method_id), r.Name), 35)

				if m.cursor.colIndex == i && m.cursor.reqIndex != nil && *m.cursor.reqIndex == j {
					subTree.Child(selectedStyle.Render(label))
				} else {
					subTree.Child(normalStyle.Render(label))
				}
			}
		}
	}

	containerBoxStyle := lipgloss.NewStyle().
		Height(HEIGHT).
		Width(WIDTH).
		Padding(PADDING).
		Border(lipgloss.ThickBorder())

	return containerBoxStyle.Render(t.String())
}
