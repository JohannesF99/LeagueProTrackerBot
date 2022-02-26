package loader

import (
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"sort"
	"time"
)

func GetRankedDataForAllPlayers(teams *[]model.Team) []model.Player {
	getUpdatedLeaguePoints(teams)
	UpdatePlayerWatchList(*teams)
	return getSortedPlayerList(*teams)
}

func getUpdatedLeaguePoints(teams *[]model.Team) []int {
	var diff []int
	for i, team := range *teams {
		for j, player := range team.Players {
			leagueEntry := lol.GetPlayerRank(player.PuuId, "RANKED_SOLO_5x5")
			(*teams)[i].Players[j].LpDiff = leagueEntry.LeaguePoints - player.Lp
			(*teams)[i].Players[j].Lp = leagueEntry.LeaguePoints
			time.Sleep(1 * time.Second)
		}
	}
	return diff
}

func getSortedPlayerList(teams []model.Team) []model.Player {
	var playerList []model.Player
	for _, team := range teams {
		for _, player := range team.Players {
			playerList = append(playerList, player)
		}
	}
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Lp > playerList[j].Lp
	})
	return playerList
}
