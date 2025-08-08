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
		log.Println("parsed appliacation state:", data)
	}
}
