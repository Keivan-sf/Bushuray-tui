package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleServerNotifs(msg sharedtypes.ServerNotification, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sharedtypes.ApplicationState:
		return applyApplicationState(msg, m)
	case sharedtypes.ProfileUpdated:
		log.Println("reached notif handler")
		return applyProfileUpdated(msg, m)
	}
	return m, nil
}
