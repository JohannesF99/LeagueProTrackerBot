package model

type InfoDTO struct {
	GameCreation       int64            `json:"gameCreation"`
	GameDuration       int              `json:"gameDuration"`
	GameEndTimestamp   int64            `json:"gameEndTimestamp"`
	GameId             int64            `json:"gameId"`
	GameMode           string           `json:"gameMode"`
	GameName           string           `json:"gameName"`
	GameStartTimestamp int64            `json:"gameStartTimestamp"`
	GameType           string           `json:"gameType"`
	GameVersion        string           `json:"gameVersion"`
	MapId              int              `json:"mapId"`
	Participants       []ParticipantDTO `json:"participants"`
	PlatformId         string           `json:"platformId"`
	QueueId            int              `json:"queueId"`
	Teams              []TeamDTO        `json:"teams"`
	TournamentCode     string           `json:"tournamentCode"`
}
