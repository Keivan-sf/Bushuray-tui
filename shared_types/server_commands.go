package sharedtypes

type GetApplicationStateData struct{}

type ConnectData struct {
	Profile ProfileID `json:"profile"`
}

type TestProfileData struct {
	Profile ProfileID `json:"profile"`
}

type DisconnectData struct{}

type AddProfilesData struct {
	Uris    string `json:"uris"`
	GroupId int    `json:"group_id"`
}

type DeleteProfilesData struct {
	Profiles []ProfileID `json:"profiles"`
}

type AddGroupData struct {
	Name            string `json:"name"`
	SubscriptionUrl string `json:"subscription_url"`
}

type ProfileID struct {
	Id      int `json:"id"`
	GroupId int `json:"group_id"`
}

type Message[T any] struct {
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
