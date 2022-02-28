package utils

import (
	"encoding/json"
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/api/lol"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"io/ioutil"
	"os"
	"time"
)

func LoadRegionWatchlist() model.Region {
	var region model.Region
	regionAsJson, _ := os.Open("./config/player_watchlist.json")
	defer regionAsJson.Close()
	byteValue, _ := ioutil.ReadAll(regionAsJson)
	err := json.Unmarshal(byteValue, &region)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return region
}

func UpdatePuuidForAllTeams(region model.Region) {
	for i, s := range region.Teams {
		for j, si := range s.Players {
			region.Teams[i].Players[j].PuuId = lol.GetPlayerPUUIDBySummonerID(si.SummonerName)
		}
		time.Sleep(2 * time.Second)
	}
	UpdatePlayerWatchList(region)
}

func UpdatePlayerWatchList(region model.Region) {
	file, _ := json.MarshalIndent(region, "", " ")
	_ = ioutil.WriteFile("./config/player_watchlist.json", file, 0644)
}
