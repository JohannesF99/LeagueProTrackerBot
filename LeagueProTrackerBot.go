package main

import (
	"fmt"
	"github.com/JohannesF99/LeagueProTrackerBot/api/twitter"
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
	println(twitter.Tweet(generateTextBody(loader.GetRankedDataForAllPlayers(&teams))))
}

func generateTextBody(ranks []model.Player) string {
	date := time.Now()
	body := "💫 @PrimeLeague SoloQ " + date.Format("02.01.2006") + " 💫\n\n"
	tabs := getTabs(ranks)
	var diffEmoji string
	for i := 0; i < 5; i++ {
		switch {
		case ranks[i].LpDiff > 0:
			diffEmoji = "+"
		case ranks[i].LpDiff == 0:
			diffEmoji = "±"
		default:
			diffEmoji = ""
		}
		body += emoji[i] + " " + ranks[i].PlayerName + tabs[i] + "(" + strconv.Itoa(ranks[i].Lp) + "LP | " + diffEmoji + strconv.Itoa(ranks[i].LpDiff) + "LP)\n"
	}
	return body
}

func getTabs(ranks []model.Player) []string {
	maxLength := 0
	var tabs []string
	for _, s := range ranks {
		if maxLength < len(s.PlayerName) {
			maxLength = len(s.PlayerName)
		}
	}
	for _, s := range ranks {
		c := maxLength - len(s.PlayerName)
		tabs = append(tabs, strings.Repeat(" ", c))
	}
	return tabs
}
