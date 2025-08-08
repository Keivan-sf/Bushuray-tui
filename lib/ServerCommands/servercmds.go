package servercmds

import (
	connection "bushuray-tui/lib/Connection"
	"bushuray-tui/shared_types"
	"encoding/json"
	"log"
)

type ServerCmds struct {
	connection *connection.ConnectionHandler
}

var sc ServerCmds

func Init(connection *connection.ConnectionHandler) {
	sc.connection = connection
}

func GetApplicationState() {
	sendCmd("get-application-state", sharedtypes.GetApplicationStateData{})
}

func sendCmd(msg string, obj any) {
	cmd := CreateJsonCommand(msg, obj)
	err := sc.connection.Send(cmd)
	if err != nil {
		log.Fatalf("failed to send command %v\n", string(cmd))
	}
}

func CreateJsonCommand(msg string, obj any) []byte {
	data := sharedtypes.Message[any]{
		Msg:  msg,
		Data: obj,
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("failed to parse json trying to send a message %v %s", data, err)
	}
	return json_data
}
