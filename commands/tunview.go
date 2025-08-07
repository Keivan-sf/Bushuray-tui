package cmds

import (
	t "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func EnterTunView() tea.Msg {
	return t.TunViewEnter{}
}

func KillCore() tea.Msg {
	return t.KillCore{}
}

func ExitTunView() tea.Msg {
	return t.TunViewExit{}
}
