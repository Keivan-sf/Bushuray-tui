package servernotifs

import (
	sharedtypes "bushuray-tui/shared_types"
	"encoding/json"
	"log"
)

func HandleNotification(msg sharedtypes.TcpMessage) {
	switch msg.Msg {
	case "application-state":
		var data sharedtypes.ApplicationState
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for application-state %v", err)
			return
		}
		ApplicationStateNotif(data)
	case "profile-updated":
		var data sharedtypes.ProfileUpdated
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profile-updated %v", err)
			return
		}
		ProfileUpdatedNotif(data)
	case "status-changed":
		var data sharedtypes.ProxyStatus
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for status-changed %v", err)
			return
		}
		StatusChangedNotif(data)
	case "profiles-added":
		var data sharedtypes.ProfilesAdded
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profiles-added %v", err)
			return
		}
		ProfilesAddedNotif(data)
	case "profiles-deleted":
		var data sharedtypes.ProfilesDeleted
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Invalid body for profiles-deleted %v", err)
			return
		}
		ProfilesDeletedNotif(data)
	}
}
