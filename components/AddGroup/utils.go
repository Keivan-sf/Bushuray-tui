package addgroup

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) adjustToNewFocus(commands []tea.Cmd) []tea.Cmd {
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			commands = append(commands, m.inputs[i].Focus())
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		}
		// Remove focused state
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}
	return commands
}

func (m *Model) reset() {
	m.focusIndex = 0
	m.inputs[0].Reset()
	m.inputs[1].Reset()
}
