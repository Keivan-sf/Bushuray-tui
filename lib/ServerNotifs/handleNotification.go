package servernotifs

import (
	"encoding/json"
	"log"

	np "github.com/Keivan-sf/Bushuray-tui/lib/NotifPublisher"
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"
)

func HandleNotification(msg sharedtypes.TcpMessage) {
	switch msg.Msg {
	case "warn":
		var data sharedtypes.Warning
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for warning %v", err)
			return
		}
		np.WarningNotif(data)
	case "application-state":
		var data sharedtypes.ApplicationState
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for application-state %v", err)
			return
		}
		np.ApplicationStateNotif(data)
	case "profile-updated":
		var data sharedtypes.ProfileUpdated
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profile-updated %v", err)
			return
		}
		np.ProfileUpdatedNotif(data)
	case "status-changed":
		var data sharedtypes.ProxyStatus
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for status-changed %v", err)
			return
		}
		np.StatusChangedNotif(data)
	case "profiles-added":
		var data sharedtypes.ProfilesAdded
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profiles-added %v", err)
			return
		}
		np.ProfilesAddedNotif(data)
	case "profiles-deleted":
		var data sharedtypes.ProfilesDeleted
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profiles-deleted %v", err)
			return
		}
		np.ProfilesDeletedNotif(data)
	case "group-added":
		var data sharedtypes.GroupAdded
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for group-added %v", err)
			return
		}
		np.GroupAddedNotif(data)
	case "group-deleted":
		var data sharedtypes.GroupDeleted
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for group-deleted %v", err)
			return
		}
		np.GroupDeletedNotif(data)
	case "subscription-updated":
		var data sharedtypes.SubscriptionUpdated
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for subscription-updated %v", err)
			return
		}
		np.SubscriptionUpdatedNotif(data)
	case "is-root-answer":
		var data sharedtypes.IsRootAnswer
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for is-root-answer %v", err)
			return
		}
		np.IsRootAnswerNotif(data)

	case "tun-status-changed":
		var data sharedtypes.TunStatus
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for tun-status-changed %v", err)
			return
		}
		np.TunStatusChangedNotif(data)
	}
}
