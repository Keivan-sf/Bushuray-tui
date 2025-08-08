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
	}
}
