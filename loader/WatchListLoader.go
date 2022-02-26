package loader

import (
	"encoding/json"
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"io/ioutil"
	"os"
	"time"
)

func LoadPrimeleagueWatchlist() []model.Team {
	var teams []model.Team
	jsonFile, _ := os.Open("./config/player_watchlist.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &teams)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return teams
}

func UpdatePuuidForAllTeams(teams *[]model.Team) {
	for i, s := range *teams {
		for j, si := range s.Players {
			(*teams)[i].Players[j].PuuId = lol.GetPlayerPUUIDBySummonerID(si.SummonerName)
		}
		time.Sleep(2 * time.Second)
	}
	UpdatePlayerWatchList(*teams)
}

func UpdatePlayerWatchList(teams []model.Team) {
	file, _ := json.MarshalIndent(teams, "", " ")
	_ = ioutil.WriteFile("./config/player_watchlist.json", file, 0644)
}
