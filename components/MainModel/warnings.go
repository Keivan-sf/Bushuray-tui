package mainmodel

import (
	notif_publisher "bushuray-tui/lib/NotifPublisher"
	sharedtypes "bushuray-tui/shared_types"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func applyWarning(msg sharedtypes.Warning, m Model) (tea.Model, tea.Cmd) {
	if msg.Key == "enable-tun-failed" {
		m.Tabs.TunStatus = "disconnected"
	}
	if msg.Key == "update-subscription-failed" {
		m.Tabs.Warning = msg.Content
		m.Tabs.LastWarningTime = time.Now()
		go func() {
			time.Sleep(time.Second * 4)
			notif_publisher.ClearWarningsNotif(sharedtypes.ClearWarnings{})
		}()
		return m, nil
	}
	log.Println("[server warning]", msg.Key, msg.Content)
	return m, nil
}
