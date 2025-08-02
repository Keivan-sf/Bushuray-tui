package tabs

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Model struct {
	Id        string
	Children  []TabView
	ActiveTap int
}

func (m Model) View() string {
	active := m.Children[m.ActiveTap]
	var tab_titles []string
	for i, child := range m.Children {
		if i == m.ActiveTap {
			tab_titles = append(tab_titles, zone.Mark(m.Id+strconv.Itoa(i), renderActiveTitle(child.Title)))
		} else {
			tab_titles = append(tab_titles, zone.Mark(m.Id+strconv.Itoa(i), renderTitle(child.Title)))
		}
	}
	tab_row := lipgloss.JoinHorizontal(lipgloss.Top, tab_titles...)
	return lipgloss.JoinVertical(lipgloss.Top, tab_row, active.View())
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	active := m.Children[m.ActiveTap]
	var cmd tea.Cmd
	m.Children[m.ActiveTap], cmd = active.Update(msg)
	return m, cmd
}

func (m Model) SetWH(width int, height int) Model {
	m.Children[m.ActiveTap] = m.Children[m.ActiveTap].SetWH(width, height)
	return m
}
