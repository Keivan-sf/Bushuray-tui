package cmds

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	t "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func SubmitGroup(name string, url string) tea.Cmd {
	return func() tea.Msg {
		servercmds.AddGroup(name, url)
		return t.AddGroupExit{}
	}
}

func ExitAddGroupView() tea.Msg {
	return t.AddGroupExit{}
}

func EnterAddGroupView() tea.Msg {
	return t.AddGroupEnter{}
}
