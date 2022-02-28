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
	regionAsJson, _ := os.Open(GetFilename())
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
		time.Sleep(1 * time.Second)
	}
	updatePlayerWatchList(region)
}

func updatePlayerWatchList(region model.Region) {
	file, _ := json.MarshalIndent(region, "", " ")
	_ = ioutil.WriteFile(GetFilename(), file, 0644)
}

func GetFilename() string {
	return "./config/" + GetRegionName() + "_watchlist.json"
}
