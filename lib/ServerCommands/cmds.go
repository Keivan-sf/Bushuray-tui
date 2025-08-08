package servercmds

import (
	sharedtypes "bushuray-tui/shared_types"
)

func GetApplicationState() {
	sendCmd("get-application-state", sharedtypes.GetApplicationStateData{})
}

func Connect(group_id int, profile_id int) {
	sendCmd("connect", sharedtypes.ConnectData{Profile: sharedtypes.ProfileID{Id: profile_id, GroupId: group_id}})
}

func Test(group_id int, profile_id int) {
	sendCmd("test-profile", sharedtypes.TestProfileData{Profile: sharedtypes.ProfileID{Id: profile_id, GroupId: group_id}})
}
