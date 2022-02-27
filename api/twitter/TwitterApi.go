package twitter

import (
	"bytes"
	"encoding/json"
	"github.com/JohannesF99/LeagueProTrackerBot/config"
	"github.com/JohannesF99/LeagueProTrackerBot/model"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
)

const url = "https://api.twitter.com/2/tweets"

func Tweet(message []string) ([]string, error) {
	client := getClient(config.GetTwitterProperties())
	var resp []string
	//Send first tweet
	body, err := generateTweetRequestBody(message[0])
	if err != nil {
		return []string{}, err
	}
	post, err := client.Post(url, "application/json", body)
	if err != nil {
		return []string{}, err
	}
	defer post.Body.Close()
	resp = append(resp, post.Status)

	//parse Response to object
	respBody, _ := ioutil.ReadAll(post.Body)
	var response model.TweetResponseBody
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return []string{}, err
	}

	//send second tweet
	body2, err := generateReplyRequestBody(message[1], response.Data.Id)
	if err != nil {
		return []string{}, err
	}
	post2, err := client.Post(url, "application/json", body2)
	if err != nil {
		return []string{}, err
	}
	resp = append(resp, post2.Status)
	return resp, nil
}

func generateTweetRequestBody(message string) (*bytes.Buffer, error) {
	body, err := json.Marshal(map[string]string{
		"text": message,
	})
	if err != nil {
		return &bytes.Buffer{}, err
	}
	return bytes.NewBuffer(body), nil
}

func generateReplyRequestBody(message string, id string) (*bytes.Buffer, error) {
	body, err := json.Marshal(model.TwitterReplyRequestBody{
		Text: message,
		Reply: struct {
			InReplyToTweetId string `json:"in_reply_to_tweet_id"`
		}{InReplyToTweetId: id},
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
