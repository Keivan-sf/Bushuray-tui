package tabs

import (
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderWarnings() string {
	color := "#df8e1d"
	if m.WarningMode == "success" {
		color = "#40a02b"
	} else if m.WarningMode == "fatal" {
		color = "#e64553"
	}
	return lipgloss.NewStyle().MarginTop(1).Height(1).Background(global.GetBgColor()).Foreground(lipgloss.Color(color)).Render(m.Warning)
}
