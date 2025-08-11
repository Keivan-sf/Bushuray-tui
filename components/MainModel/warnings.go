package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func applyWarning(msg sharedtypes.Warning, m Model) (tea.Model, tea.Cmd) {
	if msg.Key == "enable-tun-failed" {
		m.Tabs.TunStatus = "disconnected"
	}
	log.Println("[server warning]", msg.Key, msg.Content)
	return m, nil
}
