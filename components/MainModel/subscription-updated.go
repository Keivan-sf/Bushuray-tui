package mainmodel

import (
	"time"

	list "github.com/Keivan-sf/Bushuray-tui/components/List"
	notif_publisher "github.com/Keivan-sf/Bushuray-tui/lib/NotifPublisher"
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applySubscriptionUpdated(msg sharedtypes.SubscriptionUpdated, m Model) (tea.Model, tea.Cmd) {
	tid := findGroupTab(msg.GroupId, m)

	if tid == -1 {
		return m, nil
	}

	items := []list.ListItem{}

	for _, profile := range msg.Profiles {
		child := list.ListItem{
			Name:       profile.Name,
			ProfileId:  profile.Id,
			Protocol:   convertProtocolForDisplay(profile.Protocol),
			TestResult: profile.TestResult,
			Uri:        profile.Uri,
		}
		items = append(items, child)
	}

	m.Tabs.Children[tid].Content.Items = items
	if m.Tabs.Children[tid].Content.Primary != -1 {
		m.Tabs.Children[tid].Content.Primary = 0
	}
	m.Tabs.Children[tid].Content.ResetCursor()
	m.Tabs.Warning = "Subscription updated"
	m.Tabs.WarningMode = "success"
	m.Tabs.LastWarningTime = time.Now()
	go func() {
		time.Sleep(time.Second * 4)
		notif_publisher.ClearWarningsNotif(sharedtypes.ClearWarnings{})
	}()
	return m, nil
}
