package twitter

import (
	"bytes"
	"encoding/json"
	"github.com/JohannesF99/LeagueProTrackerBot/config"
	"github.com/dghubble/oauth1"
	"net/http"
)

const url = "https://api.twitter.com/2/tweets"

func Tweet(message string) string {
	client := getClient(config.GetTwitterProperties())
	body, err := generateRequestBody(message)
	if err != nil {
		return err.Error()
	}
	post, err := client.Post(url, "application/json", body)
	if err != nil {
		return err.Error()
	}
	return post.Status
}

func generateRequestBody(message string) (*bytes.Buffer, error) {
	body, err := json.Marshal(map[string]string{
		"text": message,
	})
	if err != nil {
		return &bytes.Buffer{}, err
	}
	return bytes.NewBuffer(body), nil
}

func getClient(creds config.Credentials) *http.Client {
	oauthConfig := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.Token, creds.TokenSecret)
	return oauthConfig.Client(oauth1.NoContext, token)
}
