package main

import (
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/api/twitter"
	"github.com/JohannesF99/LeagueProTrackerBot/utils"
	"strings"
)

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	region := utils.LoadRegionWatchlist()
	if utils.ShouldUpdatePuuid() {
		utils.UpdatePuuidForAllTeams(&region)
	} else {
		tweet, _ := twitter.Tweet(utils.GenerateTextBody(region, utils.GetRankedDataForAllPlayers(region), utils.GetTweetCount()))
		println(strings.Join(tweet, "\n"))
		//println(strings.Join(utils.GenerateTextBody(region, utils.GetRankedDataForAllPlayers(region), utils.GetTweetCount()), "\n\n"))
	}
}
