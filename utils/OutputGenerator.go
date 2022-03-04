package utils

import (
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

func GenerateTextBody(region model.Region, ranking []model.Player, pages int) []string {
	date := time.Now()
	bodyList := generateLadderAsString(ranking, pages)
	bodyList[0] = "💫 " + region.LeagueName + " SoloQ Ranking " + date.Format("02.01.") + " 💫\n\n" + bodyList[0] + "\n" + region.Hashtag
	return bodyList
}

func GenerateChampionsTextBody(region model.Region, ranking []model.ChampionMap, pages int) []string {
	date := time.Now()
	bodyList := generateChampionLadderAsString(ranking, pages)
	bodyList[0] = "🔥 " + region.LeagueName + " SoloQ Most Played " + date.Format("02.01.") + " 🔥\n\n" + bodyList[0] + "\n" + region.Hashtag
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

func generateChampionLadderAsString(ranks []model.ChampionMap, pages int) []string {
	var bodyList []string
	for i := 0; i < pages; i++ {
		var body string
		for j := i * 5; j < (i*5)+5; j++ {
			body += emoji[j] + " " +
				ranks[j].ChampionName + "   (" + strconv.Itoa(ranks[j].TimesPlayed) + " Games | " +
				strconv.FormatFloat(getPresencePercentage(ranks[j].TimesPlayed), 'f', 2, 32) + "%)\n"
		}
		bodyList = append(bodyList, body)
	}
	return bodyList
}

func getPresencePercentage(timesPlayed int) float64 {
	return (float64(timesPlayed) / 1000) * 100
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
