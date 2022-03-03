package model

type MetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}
