package model

type TwitterResponseBody struct {
	Data struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}
