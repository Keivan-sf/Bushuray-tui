package servernotifs

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

type ServerNotifs struct {
	p *tea.Program
}

var sn ServerNotifs

func Init(p *tea.Program) {
	sn.p = p
}

func ApplicationStateNotif(state sharedtypes.ApplicationState) {
	sn.p.Send(state)
}
