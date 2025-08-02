package tabs

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Model struct {
	Id        string
	Children  []TabView
	ActiveTap int
	Width     int
	Height    int
}

func (m Model) View() string {
	active := m.Children[m.ActiveTap]
	var tab_titles []string
	titles_len := 0
	for i, child := range m.Children {
		if i == m.ActiveTap {
			title_box := zone.Mark(m.Id+strconv.Itoa(i), renderActiveTitle(child.Title))
			titles_len += lipgloss.Width(title_box)
			tab_titles = append(tab_titles, title_box)
		} else {
			title_box := zone.Mark(m.Id+strconv.Itoa(i), renderTitle(child.Title))
			titles_len += lipgloss.Width(title_box)
			tab_titles = append(tab_titles, title_box)
		}
	}
	extra_line_w := m.Width - titles_len
	if extra_line_w > 0 {
		extra_line_str := strings.Repeat(" ", extra_line_w)
		extra_line := lipgloss.NewStyle().
			Border(lipgloss.Border{Bottom: "─", BottomLeft: "─"}).
			Render(extra_line_str)
		tab_titles = append(tab_titles, extra_line)
	}
	tab_row := lipgloss.JoinHorizontal(lipgloss.Top, tab_titles...)
	return lipgloss.JoinVertical(lipgloss.Top, tab_row, active.View())
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "ctrl+pgdown":
			if m.ActiveTap < len(m.Children)-1 {
				m.ActiveTap++
			} else {
				m.ActiveTap = 0
			}
			return m, nil
		case "shift+tab", "ctrl+pgup":
			if m.ActiveTap > 0 {
				m.ActiveTap--
			} else {
				m.ActiveTap = len(m.Children) - 1
			}
			return m, nil
		}
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonLeft:
			for i := range m.Children {
				if zone.Get(m.Id + strconv.Itoa(i)).InBounds(msg) {
					m.ActiveTap = i
					return m, nil
				}
			}
		}
	}
	var cmds []tea.Cmd
	for i, child := range m.Children {
		var cmd tea.Cmd
		m.Children[i], cmd = child.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = height
	for i, child := range m.Children {
		m.Children[i] = child.SetWH(width, height)
	}
	return m
}
