package model

import "strconv"

type SummonerDTO struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (summoner SummonerDTO) ToString() string {
	return "SummonerDTO{\n" +
		"\tid:\t\t" + summoner.Id + "\n" +
		"\taccount_id:\t" + summoner.AccountId + "\n" +
		"\tpuuid:\t\t" + summoner.Puuid + "\n" +
		"\tname:\t\t" + summoner.Name + "\n" +
		"\tprofileIconId:\t" + strconv.Itoa(summoner.ProfileIconId) + "\n" +
		"\trevisionDate:\t" + strconv.Itoa(summoner.RevisionDate) + "\n" +
		"\tsummonerLevel:\t" + strconv.Itoa(summoner.SummonerLevel) + "\n" +
		"}"
}
