package tabs

import "github.com/charmbracelet/lipgloss"

func (m Model) renderAppTitle() string {
	title := "Bushuray-tui"
	version := "v0.1.0"
	title_text := lipgloss.NewStyle().Foreground(lipgloss.Color("#f5c2e7")).Render(title)
	separtor := lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd")).Render(" • ")
	version_text := lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd")).Render(version)
	content := lipgloss.JoinHorizontal(lipgloss.Top, title_text, separtor, version_text)
	container := lipgloss.NewStyle().Width(m.Width).Height(3).MaxHeight(3).Padding(1, 0).Align(lipgloss.Center).Render(content)
	return container
}
