package cmds

import (
	t "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func ExitUpdateProfileView() tea.Msg {
	return t.UpdateProfileExit{}
}

func EnterUpdateProfileView() tea.Msg {
	return t.UpdateProfileEnter{}
}
