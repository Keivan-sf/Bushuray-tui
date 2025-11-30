package cmds

import (
	servercmds "github.com/Keivan-sf/Bushuray-tui/lib/ServerCommands"
	t "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func EnterTunView() tea.Msg {
	servercmds.IsRoot()
	return nil
}

func KillCore() tea.Msg {
	return t.KillCore{}
}

func ExitTunView() tea.Msg {
	return t.TunViewExit{}
}
