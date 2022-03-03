package utils

import (
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"sort"
	"time"
)

func GetMostPlayedChampions(region model.Region) []model.ChampionMap {
	mostChampionsPlayed := make(map[string]int)
	for _, s := range region.Teams {
		for _, si := range s.Players {
			for _, sj := range lol.GetPlayerMatchIdsByPUUID(si.PuuId) {
				details := lol.GetMatchDetailsByMatchId(sj)
				championName := details.Info.Participants[getPlayerIndex(si.PuuId, details.Metadata.Participants)].ChampionName
				mostChampionsPlayed[championName]++
				println(championName)
				time.Sleep(1200 * time.Millisecond)
			}
		}
	}
	return parseMapToSortedArray(mostChampionsPlayed)
}

func getPlayerIndex(puuid string, participants []string) int {
	for i, s := range participants {
		if s == puuid {
			return i
		}
	}
	return 11
}

func parseMapToSortedArray(championsMap map[string]int) []model.ChampionMap {
	var championMapAsArray []model.ChampionMap
	for k, v := range championsMap {
		championMapAsArray = append(championMapAsArray, model.ChampionMap{ChampionName: k, TimesPlayed: v})
	}
	sort.Slice(championMapAsArray, func(i, j int) bool {
		return championMapAsArray[i].TimesPlayed > championMapAsArray[j].TimesPlayed
	})
	return championMapAsArray
}
