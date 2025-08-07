package tabs

import (
	cmds "bushuray-tui/commands"
	"strconv"

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
	viewStart int
	viewEnd   int
}

func (m Model) View() string {
	active := m.Children[m.ActiveTap]
	var tab_titles []string
	titles_len := 0
	for i := m.viewStart; i < len(m.Children); i++ {
		child := m.Children[i]
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
		tab_titles = append(tab_titles, renderTabLine(extra_line_w))
	}
	tab_row := zone.Mark(m.Id+"tabline", lipgloss.JoinHorizontal(lipgloss.Top, tab_titles...))
	return lipgloss.JoinVertical(lipgloss.Top, m.renderAppTitle(), m.renderHelp(), tab_row, active.View())
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
			m.adjustView()
			return m, nil
		case "shift+tab", "ctrl+pgup":
			if m.ActiveTap > 0 {
				m.ActiveTap--
			} else {
				m.ActiveTap = len(m.Children) - 1
			}
			m.adjustView()
			return m, nil
		case "a":
			return m, cmds.EnterAddGroupView
		case "v":
			return m, cmds.EnterTunView
		}

		var cmd tea.Cmd
		last_primary := m.Children[m.ActiveTap].Content.Primary
		m.Children[m.ActiveTap], cmd = m.Children[m.ActiveTap].Update(msg)
		if last_primary != m.Children[m.ActiveTap].Content.Primary {
			// log.Panicf("%d: primary changed from %d to %d", m.ActiveTap,m.)
			for i := range m.Children {
				if i == m.ActiveTap {
					continue
				}
				m.Children[i].Content.Primary = -1
			}
		}

		return m, cmd
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonWheelDown:
			if !zone.Get(m.Id + "tabline").InBounds(msg) {
				break
			}
			if m.ActiveTap < len(m.Children)-1 {
				m.ActiveTap++
			}
			m.adjustView()
		case tea.MouseButtonWheelUp:
			if !zone.Get(m.Id + "tabline").InBounds(msg) {
				break
			}
			if m.ActiveTap > 0 {
				m.ActiveTap--
			}
			m.adjustView()
		case tea.MouseButtonLeft:
			if msg.Action != tea.MouseActionRelease {
				break
			}
			for i := range m.Children {
				if zone.Get(m.Id + strconv.Itoa(i)).InBounds(msg) {
					m.ActiveTap = i
					m.adjustView()
					return m, nil
				}
			}
		}

		var cmd tea.Cmd
		m.Children[m.ActiveTap], cmd = m.Children[m.ActiveTap].Update(msg)
		return m, cmd
	}
	// var cmds []tea.Cmd
	// for i, child := range m.Children {
	// 	var cmd tea.Cmd
	// 	m.Children[i], cmd = child.Update(msg)
	// 	cmds = append(cmds, cmd)
	// }
	//
	// return m, tea.Batch(cmds...)
	return m, nil
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = height
	for i, child := range m.Children {
		m.Children[i] = child.SetWH(width, height)
	}
	m.adjustToDimentions()
	return m
}
