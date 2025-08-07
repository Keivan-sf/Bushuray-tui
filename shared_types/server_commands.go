package sharedtypes

type GetApplicationStateData struct{}

type Message[T any] struct {
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
