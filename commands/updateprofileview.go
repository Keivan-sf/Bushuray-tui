package cmds

import (
	t "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func ExitUpdateProfileView() tea.Msg {
	return t.UpdateProfileExit{}
}

func EnterUpdateProfileView() tea.Msg {
	return t.UpdateProfileEnter{}
}
