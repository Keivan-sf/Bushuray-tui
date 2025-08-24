package tabs

import (
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderWarnings() string {
	return lipgloss.NewStyle().MarginTop(1).Height(1).Background(global.GetBgColor()).Foreground(lipgloss.Color("#df8e1d")).Render(m.Warning)
}
