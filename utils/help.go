package utils

import "github.com/charmbracelet/lipgloss"

func GenHelp(keys []string, helps []string) string {
	separtor := " • "
	// keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd"))
	keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7c7f93"))

	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4c4f69"))
	s := ""
	for i := range keys {
		s += keyStyle.Render(keys[i])
		s += " " + helpStyle.Render(helps[i])
		if i+1 != len(keys) {
			s += separtor
		}
	}
	return s
}
