package utils

import (
	"flag"
	"os"
	"strings"
)

var (
	regionName  = flag.String("region", "primeleague", "the region, which should be generated\n")
	updatePuuid = flag.Bool("update-puuid", false, "updates the puuid based on the given summoner name")
	mostPlayed  = flag.Bool("most-played", false, "will generate a list of most played champion instead of the SoloQ ranking")
	tweetCount  = flag.Int("tweets", 2, "how many tweets should be send (each one containing 5 players)\nmin=1\nmax=10\n")
)

func GetTweetCount() int {
	flag.Parse()
	if *tweetCount <= 0 {
		return 1
	}
	if *tweetCount > 10 {
		return 10
	}
	return *tweetCount
}

func GetRegionName() string {
	flag.Parse()
	if strings.Trim(*regionName, " ") == "" {
		println("Region invalid")
		os.Exit(2)
	}
	return strings.ToLower(*regionName)
}

func ShouldUpdatePuuid() bool {
	flag.Parse()
	return *updatePuuid
}

func GenerateChampionList() bool {
	flag.Parse()
	return *mostPlayed
}
