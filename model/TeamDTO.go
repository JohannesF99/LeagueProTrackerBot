package model

type TeamDTO struct {
	Bans       []BanDTO      `json:"bans"`
	Objectives ObjectivesDTO `json:"objectives"`
	TeamId     int           `json:"teamId"`
	Win        bool          `json:"win"`
}

type BanDTO struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type ObjectivesDTO struct {
	Baron      ObjectiveDTO `json:"baron"`
	Champion   ObjectiveDTO `json:"champion"`
	Dragon     ObjectiveDTO `json:"dragon"`
	Inhibitor  ObjectiveDTO `json:"inhibitor"`
	RiftHerald ObjectiveDTO `json:"riftHerald"`
	Tower      ObjectiveDTO `json:"tower"`
}

type ObjectiveDTO struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}
