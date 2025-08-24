package notif_publisher

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

type NotifPublisher struct {
	p *tea.Program
}

var np NotifPublisher

func Init(p *tea.Program) {
	np.p = p
}

func ApplicationStateNotif(state sharedtypes.ApplicationState) {
	np.p.Send(state)
}

func ProfileUpdatedNotif(data sharedtypes.ProfileUpdated) {
	np.p.Send(data)
}

func StatusChangedNotif(data sharedtypes.ProxyStatus) {
	np.p.Send(data)
}

func ProfilesAddedNotif(data sharedtypes.ProfilesAdded) {
	np.p.Send(data)
}

func ProfilesDeletedNotif(data sharedtypes.ProfilesDeleted) {
	np.p.Send(data)
}

func GroupAddedNotif(data sharedtypes.GroupAdded) {
	np.p.Send(data)
}

func GroupDeletedNotif(data sharedtypes.GroupDeleted) {
	np.p.Send(data)
}

func SubscriptionUpdatedNotif(data sharedtypes.SubscriptionUpdated) {
	np.p.Send(data)
}

func IsRootAnswerNotif(data sharedtypes.IsRootAnswer) {
	np.p.Send(data)
}

func TunStatusChangedNotif(data sharedtypes.TunStatus) {
	np.p.Send(data)
}

func WarningNotif(data sharedtypes.Warning) {
	np.p.Send(data)
}

func ClearWarningsNotif(data sharedtypes.ClearWarnings) {
	np.p.Send(data)
}
