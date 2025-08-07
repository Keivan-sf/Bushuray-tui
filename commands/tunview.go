package cmds

import (
	t "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func EnterTunView() tea.Msg {
	return t.TunViewEnter{}
}
