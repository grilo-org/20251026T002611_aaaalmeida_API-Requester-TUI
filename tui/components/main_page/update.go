package main_page

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// GENERAL NAVEGATION
	case tea.KeyMsg:
		switch msg.String() {
		// exit program
		case "ctrl+q", "esc":
			return m, tea.Quit

		// switch active component
		case "tab":
			// active component is the last
			if m.active_component_index == len(m.subcomponents)-1 {
				m.active_component_index = 0
			} else {
				m.active_component_index++
			}
		case "shift+tab":
			// active component is the first
			if m.active_component_index == 0 {
				m.active_component_index = len(m.subcomponents) - 1
			} else {
				m.active_component_index--
			}
		}

	// INITIAL CMD
	// FETCH COLLECTIONS FROM DB TO MODEL ARRAY POINTER
	case messages.LoadCollectionsMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}
		m.collections = msg.Collections

		var cmd tea.Cmd
		m.subcomponents[COLLECTION_MENU_INDEX], cmd = m.subcomponents[COLLECTION_MENU_INDEX].Update(msg)
		return m, cmd

	// INITIAL CMD
	// FETCH METHODS FROM DB TO REQUEST_HEADER_BODY
	case messages.LoadMethodsMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[REQUEST_HEADERS_INDEX], cmd = m.subcomponents[REQUEST_HEADERS_INDEX].Update(msg)
		return m, cmd

	// SEND A REQUEST FROM COLLECTION_MENU TO HEADER AND REQUEST_HEADER
	case messages.SendRequestMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var headerCmd, reqHeaderCmd tea.Cmd
		m.subcomponents[HEADER_INDEX], headerCmd = m.subcomponents[HEADER_INDEX].Update(msg)
		m.subcomponents[REQUEST_HEADERS_INDEX], reqHeaderCmd = m.subcomponents[REQUEST_HEADERS_INDEX].Update(msg)
		return m, tea.Batch(headerCmd, reqHeaderCmd)

	// SEND A REQUEST FROM HEADER TO MAIN MENU
	case messages.LoadRequestMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}
		var cmd tea.Cmd
		m.subcomponents[REQUEST_HEADERS_INDEX], cmd = m.subcomponents[REQUEST_HEADERS_INDEX].Update(msg)
		// TODO: mandar o request para request_response_box (ou não)
		return m, cmd

	// USER SEND AND LOADED A REQUEST
	case messages.SendStringMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}
		var cmd tea.Cmd
		m.subcomponents[REQUEST_RESPONSE_INDEX], cmd = m.subcomponents[REQUEST_RESPONSE_INDEX].Update(msg)
		// TODO: mandar o request para request_response_box (ou não)
		return m, cmd

	// USER PRESSED BUTTON IN REQUEST_HEADER TO SEND REQUEST
	// CallRequestCmd from RequestHeader's Button
	case messages.LoadResponseMsg:
		var cmd tea.Cmd
		m.subcomponents[REQUEST_RESPONSE_INDEX], cmd = m.subcomponents[REQUEST_RESPONSE_INDEX].Update(msg)
		return m, cmd

	// case messages.DeleteRequestMsg:
	// 	var headerCmd, menuCmd tea.Cmd
	// 	m.subcomponents[HEADER_INDEX], headerCmd = m.subcomponents[HEADER_INDEX].Update(
	// 		cmds.DeleteRequestCmd(m.context, m.collections[m.s]))
	// 	m.subcomponents[COLLECTION_MENU_INDEX], menuCmd = m.subcomponents[COLLECTION_MENU_INDEX].Update(
	// 		cmds.DeleteRequestCmd(m.context, &m.requests, m.selectedTabIndex))
	// 	return m, tea.Batch(headerCmd, menuCmd)

	case messages.UpdateRequestMsg:
		if msg.Err != nil {
			m.context.Logger.Println(msg.Err)
		}
	}

	var cmd tea.Cmd
	m.subcomponents[m.active_component_index], cmd = m.subcomponents[m.active_component_index].Update(msg)

	return m, cmd
}
