package model

type Region struct {
	LeagueName string `json:"leagueName"`
	Hashtag    string `json:"hashtag"`
	Teams      []Team `json:"teams"`
}
