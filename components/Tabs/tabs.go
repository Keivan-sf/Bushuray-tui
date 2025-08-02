package tabs

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Id        string
	Children  []TabView
	ActiveTap int
}

func (m Model) View() string {
	active := m.Children[m.ActiveTap]
	return active.View()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	log.Println("got update on tabs")
	active := m.Children[m.ActiveTap]
	var cmd tea.Cmd
	m.Children[m.ActiveTap], cmd = active.Update(msg)
	return m, cmd
}

func (m Model) SetWH(width int, height int) Model {
	m.Children[m.ActiveTap] = m.Children[m.ActiveTap].SetWH(width, height)
	return m
}
