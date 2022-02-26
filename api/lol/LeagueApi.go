package lol

import (
	"encoding/json"
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"io/ioutil"
	"net/http"
)

const (
	uri = "https://euw1.api.riotgames.com"
	key = "RGAPI-7459457b-8bc8-4e00-a452-9af6fc6094f3"
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

//ALL FUNCTION, WHICH RETURN RAW JSONS

func getPlayerAccountAsJSONByInGameName(inGameName string) string {
	api := uri + "/lol/summoner/v4/summoners/by-name/" + inGameName + "?api_key=" + key
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func getPlayerAccountAsJSONByPUUID(puuid string) string {
	api := uri + "/lol/summoner/v4/summoners/by-puuid/" + puuid + "?api_key=" + key
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func getPlayerRankAsJSONByID(id string) string {
	api := uri + "/lol/league/v4/entries/by-summoner/" + id + "?api_key=" + key
	response, err := http.Get(api)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}
