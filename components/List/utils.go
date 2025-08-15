package list

import (
	"bushuray-tui/global"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func styleTestPrimary(ping int) string {
	width := test_result_w
	if ping > 0 {
		return primary_style.Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render("OK " + strconv.Itoa(ping))
	} else if ping == -1 {
		return primary_style.Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render("FAILED")
	} else if ping == -2 {
		return primary_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return primary_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}

func styleTestUnderCursor(ping int) string {
	width := test_result_w
	if ping > 0 {
		return under_cursor_style.Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render("OK " + strconv.Itoa(ping))
	} else if ping == -1 {
		return under_cursor_style.Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render("FAILED")
	} else if ping == -2 {
		return under_cursor_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return under_cursor_style.Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}

func styleTestNormal(ping int) string {
	width := test_result_w
	if ping > 0 {
		return lipgloss.NewStyle().Background(global.GetBgColor()).Foreground(lipgloss.Color("#40a02b")).Width(width).MaxWidth(width).Render("OK " + strconv.Itoa(ping))
	} else if ping == -1 {
		return lipgloss.NewStyle().Background(global.GetBgColor()).Foreground(lipgloss.Color("#e64553")).Width(width).MaxWidth(width).Render("FAILED")
	} else if ping == -2 {
		return lipgloss.NewStyle().Background(global.GetBgColor()).Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("TESTING")
	} else {
		return lipgloss.NewStyle().Background(global.GetBgColor()).Foreground(lipgloss.Color("#df8e1d")).Width(width).MaxWidth(width).Render("")
	}
}

func (l *Model) ResetCursor() {
	l.cursor = 0
	l.adjustOffsetForCursor()
}

func (l *Model) MoveCursorToLastItem() {
	l.cursor = len(l.Items) - 1
	l.adjustOffsetForCursor()
}

func (l *Model) JumpToPrimary() {
	if l.Primary == -1 {
		return
	}
	l.cursor = l.Primary
	l.adjustOffsetForCursor()
}

func (l *Model) ResolveInvalidCursor() {
	items_len := len(l.Items)
	if items_len == 0 {
		l.cursor = 0
		l.adjustOffsetForCursor()
	} else if l.cursor >= items_len {
		l.cursor = items_len - 1
		l.adjustOffsetForCursor()
	} else if l.cursor < 0 {
		l.cursor = 0
		l.adjustOffsetForCursor()
	}
}

func (l *Model) adjustOffsetForCursor() {
	if l.cursor < l.offset {
		l.offset = l.cursor
	} else if l.cursor >= l.offset+l.Height {
		l.offset = l.cursor - l.Height + 1
	}
}
func (l *Model) adjustCursorForOffset() {
	if l.offset > l.cursor {
		l.cursor = l.offset
	} else if l.offset+l.Height-1 < l.cursor {
		l.cursor = l.offset + l.Height - 1
	}
}
