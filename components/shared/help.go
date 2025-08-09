package shared

import (
	"bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

var secondary_style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4c4f69"))
var primary_style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))

func GenHelp(keys []string, helps []string) string {
	separtor := secondary_style.Background(global.GetBgColor()).Render(" • ")
	var elements []string
	for i := range keys {
		key := primary_style.Background(global.GetBgColor()).Render(keys[i])
		help := secondary_style.Background(global.GetBgColor()).Render(" " + helps[i])
		if i+1 != len(keys) {
			elements = append(elements, key, help, separtor)
		} else {
			elements = append(elements, key, help)
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, elements...)
}
