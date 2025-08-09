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

func Disconnect() {
	sendCmd("disconnect", sharedtypes.DisconnectData{})
}

func Test(group_id int, profile_id int) {
	sendCmd("test-profile", sharedtypes.TestProfileData{Profile: sharedtypes.ProfileID{Id: profile_id, GroupId: group_id}})
}

func AddProfiles(uris string, gid int) {
	sendCmd("add-profiles", sharedtypes.AddProfilesData{Uris: uris, GroupId: gid})
}

func DeleteProfiles(profiles []sharedtypes.ProfileID) {
	sendCmd("delete-profiles", sharedtypes.DeleteProfilesData{Profiles: profiles})
}

func AddGroup(name string, subscription_url string) {
	sendCmd("add-group", sharedtypes.AddGroupData{Name: name, SubscriptionUrl: subscription_url})
}

func DeleteGroup(gid int) {
	sendCmd("delete-group", sharedtypes.DeleteGroupData{Id: gid})
}

func UpdateSubscription(gid int) {
	sendCmd("update-subscription", sharedtypes.UpdateSubscriptionData{GroupId: gid})
}

func IsRoot() {
	sendCmd("is-root", sharedtypes.IsRootData{})
}

func Die() {
	sendCmd("die", sharedtypes.DieData{})
}
