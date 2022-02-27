package main

import (
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/loader"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"strconv"
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
	println(generateTextBody(loader.GetRankedDataForAllPlayers(&teams)))
}

func generateTextBody(ranks []model.Player) string {
	date := time.Now()
	body := "💫 @PrimeLeague SoloQ " + date.Format("02.01.2006") + " 💫\n\n"
	body += generateLadderAsString(ranks, 1)
	body += "\n#StraussPrimeLeague "
	return body
}

func generateLadderAsString(ranks []model.Player, page int) string {
	var body string
	diffEmoji := ""
	for i := page * 5; i < (page*5)+5; i++ {
		if ranks[i].LpDiff > 0 {
			diffEmoji = "+"
		}
		if ranks[i].LpDiff == 0 {
			diffEmoji = "±"
		}
		body += emoji[i] + " " +
			ranks[i].PlayerName + " (" + strconv.Itoa(ranks[i].Lp) + "LP | " +
			diffEmoji + strconv.Itoa(ranks[i].LpDiff) + "LP)\n"
	}
	return body
}
