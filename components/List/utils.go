package list

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func styleTestPrimary(ping int) string {
	width := test_result_w
	if ping > 0 {
		return primary_style.Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -1 {
		return primary_style.Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -2 {
		return primary_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return primary_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}

func styleTestUnderCursor(ping int) string {
	width := test_result_w
	if ping > 0 {
		return under_cursor_style.Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -1 {
		return under_cursor_style.Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -2 {
		return under_cursor_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return under_cursor_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}

func styleTestNormal(ping int) string {
	width := test_result_w
	if ping > 0 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -1 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render(strconv.Itoa(ping))
	} else if ping == -2 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}
