package addgroup

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#8caaee"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	grayStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70"))
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	// blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	blurredButton = grayStyle.Render("[ Submit ]")
)
