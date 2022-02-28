package main

import (
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/loader"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"strconv"
	"strings"
	"time"
)

var emoji = map[int]string{
	0: "🏅",
	1: "🥈",
	2: "🥉",
	3: "4️⃣",
	4: "5️⃣",
	5: "6️⃣",
	6: "7️⃣",
	7: "8️⃣",
	8: "9️⃣",
	9: "🔟",
}

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	region := loader.LoadRegionWatchlist()
	//loader.UpdatePuuidForAllTeams(&teams)
	//tweet, _ := twitter.Tweet(generateTextBody(loader.GetRankedDataForAllPlayers(&teams), 2))
	//println(strings.Join(tweet, "\n"))
	println(strings.Join(generateTextBody(region, loader.GetRankedDataForAllPlayers(region), 2), "\n"))
}

func generateTextBody(region model.Region, ranking []model.Player, pages int) []string {
	date := time.Now()
	bodyList := generateLadderAsString(ranking, pages)
	bodyList[0] = "💫 " + region.LeagueName + " Ranking " + date.Format("02.01.") + " 💫\n\n" + bodyList[0] + "\n" + region.Hashtag
	return bodyList
}

func generateLadderAsString(ranks []model.Player, pages int) []string {
	var bodyList []string
	for i := 0; i < pages; i++ {
		var body string
		for j := i * 5; j < (i*5)+5; j++ {
			diffEmoji := getDifferenceEmoji(ranks[j])
			division := getDivision(ranks[j])
			body += emoji[j] + " " +
				ranks[j].PlayerName + "   (" + division + strconv.Itoa(ranks[j].Lp) + "LP | " +
				diffEmoji + strconv.Itoa(ranks[j].LpDiff) + "LP)\n"
		}
		bodyList = append(bodyList, body)
	}
	return bodyList
}

func getDivision(player model.Player) string {
	if strings.Contains(player.Division, "Chall") ||
		strings.Contains(player.Division, "GM") ||
		strings.Contains(player.Division, "MA") ||
		player.Division == "" {
		return ""
	}
	return player.Division + player.SubDiv + " "
}

func getDifferenceEmoji(player model.Player) string {
	if player.LpDiff > 0 {
		return "+"
	}
	if player.LpDiff == 0 {
		return "±"
	}
	return ""
}
