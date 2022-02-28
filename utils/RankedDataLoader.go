package utils

import (
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"sort"
	"time"
)

var (
	div = map[string]int{
		"Chall": 6,
		"GM":    6,
		"MA":    6,
		"Dia":   5,
		"Plat":  4,
		"Gold":  3,
		"Sil":   2,
		"Bro":   1,
		"Iron":  0,
	}
	sdiv = map[string]int{
		"I":   3,
		"II":  2,
		"III": 1,
		"IV":  0,
	}
)

func GetRankedDataForAllPlayers(region model.Region) []model.Player {
	getUpdatedLeaguePoints(&region.Teams)
	updatePlayerWatchList(region)
	return getSortedPlayerList(region.Teams)
}

func getUpdatedLeaguePoints(teams *[]model.Team) {
	for i, team := range *teams {
		for j, player := range team.Players {
			leagueEntry := lol.GetPlayerRank(player.PuuId, "RANKED_SOLO_5x5")
			(*teams)[i].Players[j].LpDiff = leagueEntry.LeaguePoints - player.Lp
			(*teams)[i].Players[j].Lp = leagueEntry.LeaguePoints
			(*teams)[i].Players[j].Division = getDivisionTag(leagueEntry.Tier)
			(*teams)[i].Players[j].SubDiv = leagueEntry.Rank
			time.Sleep(1 * time.Second)
		}
	}
}

func getDivisionTag(tier string) string {
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
		lp1 := getTrueLeaguePoints(playerList[i].Division, playerList[i].SubDiv, playerList[i].Lp)
		lp2 := getTrueLeaguePoints(playerList[j].Division, playerList[j].SubDiv, playerList[j].Lp)
		return lp1 > lp2
	})
	return playerList
}

func getTrueLeaguePoints(division string, subDivision string, lp int) int {
	return ((div[division]*5 + sdiv[subDivision]) * 100) + lp
}
