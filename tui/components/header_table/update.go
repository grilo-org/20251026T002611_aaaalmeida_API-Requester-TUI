package header_table

import (
	cmds "api-requester/tui/commands"
	"api-requester/tui/messages"
	"api-requester/utils"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// always have atleast 1 line
	if len(m.rows) == 0 {
		k := textinput.New()
		v := textinput.New()
		m.rows = append(m.rows, HeaderRow{KeyTextInput: k, ValueTextInput: v})
	}

	switch msg := msg.(type) {
	// INITIAL CMD
	// LOAD REQUEST
	case messages.LoadRequestMsg:
		m.rows = nil // clear cache

		for k, v := range msg.Request.Headers {
			keyInput := textinput.New()
			keyInput.SetValue(k)
			valueInput := textinput.New()
			valueInput.SetValue(v)
			m.rows = append(m.rows, HeaderRow{KeyTextInput: keyInput, ValueTextInput: valueInput})
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		// change between key or value
		case ":":
			m.keyOrValue = !m.keyOrValue

			return m, cmds.HeaderTableChangedCmd(m.headerRowToMapString())

		case "up":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
			return m, nil

		case "down":
			if m.cursorIndex < len(m.rows)-1 {
				m.cursorIndex++
			}
			return m, nil

		// create line
		case "enter":
			// point to last row
			m.cursorIndex = len(m.rows) - 1
			// row := m.rows[m.cursorIndex]
			// if row.KeyTextInput.Value() == "" || row.ValueTextInput.Value() == "" {
			// 	return m, nil
			// }
			k := textinput.New()
			v := textinput.New()

			m.rows = append(m.rows, HeaderRow{KeyTextInput: k, ValueTextInput: v})
			m.cursorIndex++

			return m, cmds.HeaderTableChangedCmd(m.headerRowToMapString())

		// delete line
		case "delete":
			if len(m.rows) > 0 {
				m.rows = utils.RemoveFromArray(m.rows, m.cursorIndex)
				if m.cursorIndex >= len(m.rows) && len(m.rows) > 0 {
					m.cursorIndex = len(m.rows) - 1
				}
			}

			return m, cmds.HeaderTableChangedCmd(m.headerRowToMapString())

		// delete char from line
		case "backspace":
			row := &m.rows[m.cursorIndex]
			if m.keyOrValue {
				row.KeyTextInput.SetValue(utils.RemoveLastChar(row.KeyTextInput.Value()))
			} else {
				row.ValueTextInput.SetValue(utils.RemoveLastChar(row.ValueTextInput.Value()))
			}
			return m, nil

		// write into line
		default:
			if utils.IsValidRequestHeaderValue(msg.String()) {
				row := &m.rows[m.cursorIndex]
				if m.keyOrValue {
					row.KeyTextInput.SetValue(row.KeyTextInput.Value() + msg.String())
				} else {
					row.ValueTextInput.SetValue(row.ValueTextInput.Value() + msg.String())
				}
				return m, cmds.HeaderTableChangedCmd(m.headerRowToMapString())
			}
		}
	}

	// component lost focus, so we save changes
	if !m.isFocused {
		m.context.Logger.Println("PERDI O FOCO")
		return m, cmds.HeaderTableChangedCmd(m.headerRowToMapString())
	}

	return m, nil
}

func (m *Model) headerRowToMapString() map[string]string {
	rowsMap := make(map[string]string, len(m.rows))
	for _, v := range m.rows {
		rowsMap[v.KeyTextInput.Value()] = v.ValueTextInput.Value()
	}
	return rowsMap
}
