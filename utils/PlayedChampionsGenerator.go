package utils

import (
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"sort"
	"time"
)

func GetMostPlayedChampions(region model.Region) []model.ChampionMap {
	mostChampionsPlayed := make(map[string]model.ChampionMap)
	for _, s := range region.Teams {
		for _, si := range s.Players {
			for _, sj := range lol.GetPlayerMatchIdsByPUUID(si.PuuId) {
				details := lol.GetMatchDetailsByMatchId(sj)
				championName := getChampionName(si.PuuId, details)
				win := isWin(getPlayerIndex(si.PuuId, details), details)
				mostChampionsPlayed[championName] = model.ChampionMap{
					ChampionName: championName,
					TimesPlayed:  mostChampionsPlayed[championName].TimesPlayed + 1,
					Wins:         mostChampionsPlayed[championName].Wins + win,
				}
				println(championName)
				time.Sleep(1200 * time.Millisecond)
			}
		}
	}
	return parseMapToSortedArray(mostChampionsPlayed)
}

func getPlayerIndex(puuid string, details model.MatchDTO) int {
	for i, s := range details.Metadata.Participants {
		if s == puuid {
			return i
		}
	}
	return 11
}

func isWin(playerId int, details model.MatchDTO) int {
	if (playerId < 5 && details.Info.Teams[0].Win) || (playerId >= 5 && details.Info.Teams[1].Win) {
		return 1
	} else {
		return 0
	}
}

func getChampionName(puuId string, details model.MatchDTO) string {
	return details.Info.Participants[getPlayerIndex(puuId, details)].ChampionName
}

func parseMapToSortedArray(championsMap map[string]model.ChampionMap) []model.ChampionMap {
	var championMapAsArray []model.ChampionMap
	for _, champion := range championsMap {
		championMapAsArray = append(championMapAsArray, champion)
	}
	sort.Slice(championMapAsArray, func(i, j int) bool {
		return championMapAsArray[i].TimesPlayed > championMapAsArray[j].TimesPlayed
	})
	return championMapAsArray
}
