package collection_menu

import (
	cmd "api-requester/tui/commands"
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// captures msg type
	switch msg := msg.(type) {

	// INITIAL UPDATE
	case messages.LoadCollectionsMsg:
		m.collections = msg.Collections
		m.openCloseIndex = make([]bool, len(m.collections))
		return m, nil

	// USER PRESSED A KEY
	case tea.KeyMsg:
		// captures which key was pressed
		switch msg.String() {
		case "delete", "backspace":
			if m.cursor.reqIndex != nil {
				// delete request
				c := m.cursor
				return m, cmd.UserPressDeleteInRequestCmd(m.collections[c.colIndex].Requests[*c.reqIndex])
			} else {
				// delete collection
			}
		case "up":
			items := m.visibleItems()
			for i := range items {
				if items[i].Equal(m.cursor) {
					if i > 0 {
						m.cursor = items[i-1]
					}
					break
				}
			}

		case "down":
			items := m.visibleItems()
			for i := range items {
				if items[i].Equal(m.cursor) {
					if i < len(items)-1 {
						m.cursor = items[i+1]
					}
					break
				}
			}

		case "enter", " ":
			if m.cursor.reqIndex == nil {
				// USER PRESSED ENTER IN A COLLECTION
				selectedIndex := m.cursor.colIndex
				m.openCloseIndex[selectedIndex] = !m.openCloseIndex[selectedIndex]

				selectedCollection := m.collections[selectedIndex]

				// LAZY LOADING.
				// ONLY LOADS REQUESTS FROM DB WHEN NEEDED.
				if selectedCollection.Requests == nil {
					return m, cmd.UserPressEnterInCollectionCmd(m.context, selectedCollection.ID)
				}
			} else {
				// USER PRESSED ENTER IN A REQUEST
				// SEND SendRequestMsg TO HEADER AND REQUEST_HEADER
				req := m.collections[m.cursor.colIndex].Requests[*m.cursor.reqIndex]
				return m, cmd.SendRequestToTabCmd(req)
			}
		}

	// USER PRESSED ENTER IN A COLLECTION.
	// IMPORT ALL REQUESTS FROM THIS COLLECTION.
	case messages.LoadRequestFromCollectionMsg:
		if msg.Err != nil {
			// TODO: tratar erro
			m.context.Logger.Println(msg.Err)
			return m, nil
		}

		/*
			FIXME: tea.Update roda em thread, talvez posso setar requests direto por indice
			com o codigo abaixo otimizando o codigo, mas caso alguma operação mude o tamanho ou
			ordem das collections isso vai resultar em bug. Testar antes de implementar mudança.
		*/
		// m.Collections[m.Cursor].Requests = msg.requests
		for i := range m.collections {
			if m.collections[i].ID == msg.Collection_id {
				m.collections[i].Requests = msg.Requests
			}
		}
	}

	return m, nil
}

// Controls tree navigation
func (m model) visibleItems() []cursor {
	var items []cursor
	for i, col := range m.collections {
		items = append(items, cursor{colIndex: i, reqIndex: nil}) // collection position
		if m.openCloseIndex[i] && col.Requests != nil {
			for j := range col.Requests {
				jCopy := j
				items = append(items, cursor{colIndex: i, reqIndex: &jCopy}) // request position
			}
		}
	}
	return items
}

// Compares value, not address
func (c cursor) Equal(other cursor) bool {
	if c.colIndex != other.colIndex {
		return false
	}
	if c.reqIndex == nil && other.reqIndex == nil {
		return true
	}
	if c.reqIndex != nil && other.reqIndex != nil {
		return *c.reqIndex == *other.reqIndex
	}
	return false
}
