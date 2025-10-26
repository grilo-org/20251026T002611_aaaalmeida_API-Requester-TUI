package request_header_box

import (
	"api-requester/domain/request"
	"api-requester/shared/focusable"
	cmds "api-requester/tui/commands"
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.selectedComponentIndex > 0 {
				m.subcomponents[m.selectedComponentIndex].Blur()
				m.selectedComponentIndex--
				m.subcomponents[m.selectedComponentIndex].Focus()
			}
			return m, nil

		case "right":
			if m.selectedComponentIndex < len(m.subcomponents)-1 {
				m.subcomponents[m.selectedComponentIndex].Blur()
				m.selectedComponentIndex++
				m.subcomponents[m.selectedComponentIndex].Focus()
			}
			return m, nil

		default:
			aux, cmd := m.subcomponents[m.selectedComponentIndex].Update(msg)
			m.subcomponents[m.selectedComponentIndex] = aux.(focusable.Focusable)
			return m, cmd
		}

	// USER PRESSED BUTTON TO SEND REQUEST
	// CallRequestCmd from Button
	case messages.ButtonPressedMsg:
		// update request then send request
		if msg.Action == ">>" && m.request != nil {
			return m, tea.Sequence(
				cmds.UpdateRequestCmd(m.context, m.request),
				cmds.CallRequestCmd(m.request))
		}

	// USER PRESSED ENTER IN HEADER_COMPONENT
	// UserPressEnterInRequestCmd from Header Component
	case messages.LoadRequestMsg:
		return m, m.handleIncommingRequest(msg.Request)

	// USER PRESSED ENTER IN COLLECTION_MENU
	// SendRequestToTabCmd from Collection_menu
	case messages.SendRequestMsg:
		return m, m.handleIncommingRequest(msg.Request)

	// INITIAL CMD
	// FETCHES METHODS FROM DB AND SEND TO SELECT COMPONENT
	case messages.LoadMethodsMsg:
		// WARNING: do not return cmd back or program will get in infinite loop
		aux, _ := m.subcomponents[SELECT_MENU_INDEX].Update(msg)
		m.subcomponents[SELECT_MENU_INDEX] = aux.(focusable.Focusable)

	// USER CHANGED REQUEST URL
	// InputChangedCmd from input_url
	case messages.InputChangedMsg:
		if m.request == nil {
			// TODO: criar erro, usuario ainda não focou um request
		}

		m.request.Url = msg.Value
		return m, cmds.UpdateRequestCmd(m.context, m.request)

	// USER CHANGED REQUEST METHOD
	// UserPressEnterInSelectCmd from Select_menu
	case messages.SendSelectMsg:
		if m.request == nil {
			// TODO: criar erro, usuario ainda não focou um request
		}

		m.request.Method_id = msg.Value.(int)
		return m, cmds.UpdateRequestCmd(m.context, m.request)

	case messages.HeaderTableChangedMsg:
		if m.request == nil {
			// TODO: criar erro, usuario ainda não focou um request
		}
		m.request.Headers = msg.Content
		m.context.Logger.Println(msg.Content)
		m.context.Logger.Println("TO SALVANDO O NOVO CABECALHO")
		return m, cmds.UpdateRequestCmd(m.context, m.request)
	}

	return m, nil
}

// Load model with outside request
func (m *Model) handleIncommingRequest(req *request.Request) tea.Cmd {
	m.request = req

	// pass url to input
	aux, inputCmd := m.subcomponents[INPUT_URL_INDEX].Update(
		messages.SendStringMsg{Value: req.Url})
	m.subcomponents[INPUT_URL_INDEX] = aux.(focusable.Focusable)

	// pass method id to select
	aux, selectCmd := m.subcomponents[SELECT_MENU_INDEX].Update(
		messages.SendNumberMsg{Value: req.Method_id})
	m.subcomponents[SELECT_MENU_INDEX] = aux.(focusable.Focusable)

	aux, tableCmd := m.subcomponents[HEADER_TABLE_INDEX].Update(
		messages.LoadRequestMsg{Request: req})
	m.subcomponents[HEADER_TABLE_INDEX] = aux.(focusable.Focusable)

	return tea.Batch(inputCmd, selectCmd, tableCmd)
}
