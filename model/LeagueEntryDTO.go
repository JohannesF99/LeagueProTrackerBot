package model

import "strconv"

type LeagueEntryDTO struct {
	LeagueId     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerId   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

func (rank LeagueEntryDTO) ToString() string {
	return "RankDTO{\n" +
		"\tLeagueId:\t" + rank.LeagueId + "\n" +
		"\tQueueType:\t" + rank.QueueType + "\n" +
		"\tTier:\t\t" + rank.Tier + "\n" +
		"\tRank:\t\t" + rank.Rank + "\n" +
		"\tSummonerId:\t" + rank.SummonerId + "\n" +
		"\tSummonerName:\t" + rank.SummonerName + "\n" +
		"\tLeaguePoints:\t" + strconv.Itoa(rank.LeaguePoints) + "\n" +
		"\tWins:\t\t" + strconv.Itoa(rank.Wins) + "\n" +
		"\tLosses:\t\t" + strconv.Itoa(rank.Losses) + "\n" +
		"\tVeteran:\t" + strconv.FormatBool(rank.Veteran) + "\n" +
		"\tInactive:\t" + strconv.FormatBool(rank.Inactive) + "\n" +
		"\tFreshBlood:\t" + strconv.FormatBool(rank.FreshBlood) + "\n" +
		"\tHotStreak:\t" + strconv.FormatBool(rank.HotStreak) + "\n" +
		"}"
}
