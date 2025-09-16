package tabs

import (
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderAppTitle() string {
	title := "Bushuray-tui"
	version := "v0.1.7"
	title_text := lipgloss.NewStyle().Foreground(lipgloss.Color("#f5c2e7")).Background(global.GetBgColor()).Render(title)
	separtor := lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd")).Background(global.GetBgColor()).Render(" • ")
	version_text := lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd")).Background(global.GetBgColor()).Render(version)
	content := lipgloss.JoinHorizontal(lipgloss.Top, title_text, separtor, version_text)
	container := lipgloss.NewStyle().Width(m.Width).Height(3).MaxHeight(3).Padding(1, 0).Background(global.GetBgColor()).Align(lipgloss.Center).Render(content)
	return container
}
