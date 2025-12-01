package addgroup

import (
	"github.com/Keivan-sf/Bushuray-tui/global"

	"github.com/charmbracelet/lipgloss"
)

var (
	placeHolderStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#313244"))
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#8caaee"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	grayStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70"))
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	// blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func focusedButton() string {
	return focusedStyle.Background(global.GetBgColor()).Render("[ Submit ]")
}

func blurredButton() string {
	return grayStyle.Background(global.GetBgColor()).Render("[ Submit ]")
}
