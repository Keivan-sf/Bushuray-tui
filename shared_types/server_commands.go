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

type ProfileID struct {
	Id      int `json:"id"`
	GroupId int `json:"group_id"`
}

type Message[T any] struct {
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
