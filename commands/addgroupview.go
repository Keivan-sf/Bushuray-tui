package cmds

import (
	t "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func ExitAddGroupView() tea.Msg {
	return t.AddGroupExit{}
}

func EnterAddGroupView() tea.Msg {
	return t.AddGroupEnter{}
}
