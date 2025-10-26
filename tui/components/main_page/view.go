package main_page

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {

	leftSide := lipgloss.JoinVertical(
		lipgloss.Top,
		m.subcomponents[SEARCH_COLLECTION_INDEX].View(),
		m.subcomponents[COLLECTION_MENU_INDEX].View(),
	)

	rightSide := lipgloss.JoinVertical(
		lipgloss.Top,
		m.subcomponents[HEADER_INDEX].View(),
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.subcomponents[REQUEST_HEADERS_INDEX].View(),
			m.subcomponents[REQUEST_RESPONSE_INDEX].View()),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftSide, rightSide)
}
