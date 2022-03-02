package lol

import (
	"encoding/json"
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/config"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"io/ioutil"
	"net/http"
)

const (
	uriEUW    = "https://euw1.api.riotgames.com"
	uriEUROPE = "https://europe.api.riotgames.com"
)

func GetPlayerRank(puuid string, queue string) model.LeagueEntryDTO {
	account := getPlayerAccountByPUUID(puuid)
	for _, rank := range getPlayerRankByID(account.Id) {
		if queue == rank.QueueType {
			return rank
		}
	}
	return model.LeagueEntryDTO{}
}

func GetPlayerPUUIDBySummonerID(inGameName string) string {
	accountJSON := getPlayerAccountAsJSONByInGameName(inGameName)
	var summoner model.SummonerDTO
	err := json.Unmarshal([]byte(accountJSON), &summoner)
	if err != nil {
		return err.Error()
	}
	return summoner.Puuid
}

func getPlayerAccountByPUUID(puuid string) model.SummonerDTO {
	accountJSON := getPlayerAccountAsJSONByPUUID(puuid)
	var summoner model.SummonerDTO
	err := json.Unmarshal([]byte(accountJSON), &summoner)
	if err != nil {
		return model.SummonerDTO{}
	}
	return summoner
}

func getPlayerRankByID(id string) []model.LeagueEntryDTO {
	rankJSON := getPlayerRankAsJSONByID(id)
	var ranks []model.LeagueEntryDTO
	err := json.Unmarshal([]byte(rankJSON), &ranks)
	if err != nil {
		return []model.LeagueEntryDTO{}
	}
	return ranks
}

func GetPlayerMatchIdsByPUUID(puuid string) []string {
	matchIdJSON := getPlayerMatchIDsAsJSONByPUUID(puuid)
	var matchIdList []string
	err := json.Unmarshal([]byte(matchIdJSON), &matchIdList)
	if err != nil {
		println(err.Error())
		return []string{}
	}
	return matchIdList
}

func GetMatchDetailsByMatchId(matchId string) model.MatchDTO {
	matchDetailsJSON := getMatchDetailsAsJSONByMatchID(matchId)
	var matchDetails model.MatchDTO
	err := json.Unmarshal([]byte(matchDetailsJSON), &matchDetails)
	if err != nil {
		println(err.Error())
		return model.MatchDTO{}
	}
	return matchDetails
}

//ALL FUNCTION, WHICH RETURN RAW JSONS

func getPlayerAccountAsJSONByInGameName(inGameName string) string {
	api := uriEUW + "/lol/summoner/v4/summoners/by-name/" + inGameName + "?api_key=" + config.GetRiotKey()
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	println(inGameName, response.Status)
	return string(data)
}

func getPlayerAccountAsJSONByPUUID(puuid string) string {
	api := uriEUW + "/lol/summoner/v4/summoners/by-puuid/" + puuid + "?api_key=" + config.GetRiotKey()
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	println(response.Status)
	return string(data)
}

func getPlayerRankAsJSONByID(id string) string {
	api := uriEUW + "/lol/league/v4/entries/by-summoner/" + id + "?api_key=" + config.GetRiotKey()
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func getPlayerMatchIDsAsJSONByPUUID(puuid string) string {
	api := uriEUROPE + "/lol/match/v5/matches/by-puuid/" + puuid + "/ids?type=ranked&start=0&count=20&api_key=" + config.GetRiotKey()
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func getMatchDetailsAsJSONByMatchID(matchId string) string {
	api := uriEUROPE + "/lol/match/v5/matches/" + matchId + "?api_key=" + config.GetRiotKey()
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}
