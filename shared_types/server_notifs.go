package sharedtypes

import "encoding/json"

type TcpMessage struct {
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type ServerNotification interface {
	IsNotification()
}

type ApplicationState struct {
	Groups           []GroupWithProfiles `json:"groups"`
	ConnectionStatus ProxyStatus         `json:"connection-status"`
	TunStatus        bool                `json:"tun-status"`
}

func (a ApplicationState) IsNotification() {}

type GroupWithProfiles struct {
	Group    Group     `json:"group"`
	Profiles []Profile `json:"profiles"`
}

type ProxyStatus struct {
	Connection string  `json:"connection"`
	Profile    Profile `json:"profile"`
}

type Profile struct {
	Id         int    `json:"id"`
	GroupId    int    `json:"group_id"`
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Uri        string `json:"uri"`
	Address    string `json:"address,omitzero"`
	Host       string `json:"host,omitzero"`
	TestResult int    `json:"test-result"`
}

type Group struct {
	Id              int    `json:"id"`
	SubscriptionUrl string `json:"subscription_url"`
	Name            string `json:"name"`
	LastId          int    `json:"last_id"`
}
