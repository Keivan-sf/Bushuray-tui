package cmds

import (
	t "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func ExitAddGroupView() tea.Msg {
	return t.AddGroupExit{}
}

func EnterAddGroupView() tea.Msg {
	return t.AddGroupEnter{}
}
