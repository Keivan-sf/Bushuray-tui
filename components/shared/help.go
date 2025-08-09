package shared

import (
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

var secondary_style = lipgloss.NewStyle().Background(global.GetBgColor()).Foreground(lipgloss.Color("#4c4f69"))
var primary_style = lipgloss.NewStyle().Background(global.GetBgColor())

func GenHelp(keys []string, helps []string) string {
	separtor := secondary_style.Render(" • ")

	s := ""
	for i := range keys {
		s += primary_style.Render(keys[i])
		s += secondary_style.Render(" " + helps[i])
		if i+1 != len(keys) {
			s += separtor
		}
	}
	return s
}
