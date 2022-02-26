package model

type Player struct {
	PlayerName   string `json:"playerName"`
	SummonerName string `json:"summonerName"`
	PuuId        string `json:"puuId"`
	Lp           int    `json:"lp"`
	LpDiff       int    `json:"lpDiff"`
}
