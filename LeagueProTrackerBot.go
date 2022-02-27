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
	teams := loader.LoadPrimeleagueWatchlist()
	//loader.UpdatePuuidForAllTeams(&teams)
	//println(twitter.Tweet(generateTextBody(loader.GetRankedDataForAllPlayers(&teams))))
	println(strings.Join(generateTextBody(loader.GetRankedDataForAllPlayers(&teams), 2), "\n\n"))
}

func generateTextBody(ranks []model.Player, pages int) []string {
	date := time.Now()
	bodyList := generateLadderAsString(ranks, pages)
	bodyList[0] = "💫 @PrimeLeague SoloQ " + date.Format("02.01.2006") + " 💫\n\n" + bodyList[0] + "\n#StraussPrimeLeague "
	return bodyList
}

func generateLadderAsString(ranks []model.Player, pages int) []string {
	var bodyList []string
	diffEmoji := ""
	for i := 0; i < pages; i++ {
		var body string
		for j := i * 5; j < (i*5)+5; j++ {
			if ranks[j].LpDiff > 0 {
				diffEmoji = "+"
			}
			if ranks[j].LpDiff == 0 {
				diffEmoji = "±"
			}
			body += emoji[j] + " " +
				ranks[j].PlayerName + " (" + strconv.Itoa(ranks[j].Lp) + "LP | " +
				diffEmoji + strconv.Itoa(ranks[j].LpDiff) + "LP)\n"
		}
		bodyList = append(bodyList, body)
	}
	return bodyList
}
