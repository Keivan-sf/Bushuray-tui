package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func applyIsRootAnswer(msg sharedtypes.IsRootAnswer, m Model) (tea.Model, tea.Cmd) {
	if !msg.IsRoot {
		m.ActiveSection = "tunview"
		return m, nil
	} else {
		log.Println("is actually root")
		return m, nil
	}
}
