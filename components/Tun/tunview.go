package tunview

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width  int
	Height int
}

func InitialModel() Model {
	return Model{}
}

func (m Model) View() string {
	return "hiiii"
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = width
	return m
}
