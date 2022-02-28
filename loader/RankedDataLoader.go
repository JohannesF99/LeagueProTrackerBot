package loader

import (
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"sort"
	"time"
)

var div = map[string]int{
	"Chall": 7,
	"GM":    7,
	"MA":    7,
	"Dia":   6,
	"Plat":  5,
	"Gold":  4,
	"Sil":   3,
	"Bro":   2,
	"Iron":  1,
}

var sdiv = map[string]int{
	"I":   5,
	"II":  4,
	"III": 3,
	"IV":  2,
	"V":   1,
}

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
			(*teams)[i].Players[j].Division = getDivision(leagueEntry.Tier)
			(*teams)[i].Players[j].SubDiv = leagueEntry.Rank
			time.Sleep(1 * time.Second)
		}
	}
	return diff
}

func getDivision(tier string) string {
	var division string
	switch tier {
	case "CHALLENGER":
		division = "Chall"
	case "GRANDMASTER":
		division = "GM"
	case "MASTER":
		division = "MA"
	case "DIAMOND":
		division = "Dia"
	case "PLATINUM":
		division = "Plat"
	case "GOLD":
		division = "Gold"
	case "SILVER":
		division = "Sil"
	case "BRONZE":
		division = "Bro"
	case "IRON":
		division = "Iron"
	default:
		division = ""
	}
	return division
}

func getSortedPlayerList(teams []model.Team) []model.Player {
	var playerList []model.Player
	for _, team := range teams {
		for _, player := range team.Players {
			playerList = append(playerList, player)
		}
	}
	sort.Slice(playerList, func(i, j int) bool {
		if div[playerList[i].Division] == div[playerList[j].Division] {
			if sdiv[playerList[i].SubDiv] == sdiv[playerList[j].SubDiv] {
				return playerList[i].Lp > playerList[j].Lp
			} else {
				return sdiv[playerList[i].SubDiv] > sdiv[playerList[j].SubDiv]
			}
		}
		return div[playerList[i].Division] > div[playerList[j].Division]
	})
	return playerList
}
