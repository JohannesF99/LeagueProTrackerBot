package model

type Team struct {
	Team    string   `json:"team"`
	Tag     string   `json:"tag"`
	Players []Player `json:"players"`
}
