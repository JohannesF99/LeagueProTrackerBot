package model

type MatchDTO struct {
	Metadata MetadataDTO `json:"metadata"`
	Info     InfoDTO     `json:"info"`
}
