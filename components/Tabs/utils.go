package tabs

import (
	"bushuray-tui/components/shared"
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderHelp() string {
	var help_text = shared.GenHelp([]string{"enter", "p", "v", "a", "t", "?"}, []string{"connect", "paste profile", "tun mode", "add group", "test", "help menu"})
	help := lipgloss.NewStyle().Width(m.Width).Height(2).MaxHeight(2).MaxWidth(m.Width).Align(lipgloss.Center).Background(global.GetBgColor()).Render(help_text)
	return help
}

func (m *Model) adjustToDimentions() {
	m.viewStart = m.ActiveTap
	m.viewEnd = m.calculateEnd(m.viewStart)
}

func (m *Model) AdjustView() {
	if m.viewStart == m.viewEnd && m.viewEnd == 0 {
		m.viewEnd = m.calculateEnd(m.viewStart)
		return
	}

	if m.ActiveTap > m.viewEnd {
		m.viewEnd = m.ActiveTap
		m.viewStart = m.calculateStart(m.viewEnd)
	} else if m.ActiveTap < m.viewStart {
		m.viewStart = m.ActiveTap
		m.viewEnd = m.calculateEnd(m.viewStart)
	}
}

func (m *Model) calculateStart(end int) int {
	start := end
	width := 0
	for i := end; i >= 0 && width <= m.Width; i-- {
		start = i
		width += len(m.Children[i].Title) + 4
	}
	if width > m.Width && end != start {
		return start + 1
	} else {
		return start
	}
}

func (m *Model) calculateEnd(start int) int {
	end := start
	width := 0
	for i := start; i < len(m.Children) && width <= m.Width; i++ {
		end = i
		width += len(m.Children[i].Title) + 4
	}
	if width > m.Width && end != start {
		return end - 1
	} else {
		return end
	}
}
