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

func Tweet(tweetList []string) ([]string, error) {
	twitterClient := getClient(config.GetTwitterProperties())
	var twitterResponseCodeList []string

	tweetId, err := sendMainTweet(tweetList[0], twitterClient, &twitterResponseCodeList)
	if err != nil {
		return twitterResponseCodeList, err
	}

	tweetList = tweetList[1:]
	for _, tweet := range tweetList {
		tweetId, err = sendReplyTweet(tweet, twitterClient, &twitterResponseCodeList, tweetId)
		if err != nil {
			return twitterResponseCodeList, err
		}
	}

	return twitterResponseCodeList, nil
}

func sendMainTweet(tweet string, twitterClient *http.Client, twitterResponseCodeList *[]string) (string, error) {
	tweetRequestBody, err := generateTweetRequestBody(tweet)
	if err != nil {
		return "", err
	}

	twitterResponse, err := twitterClient.Post(url, "application/json", tweetRequestBody)
	if err != nil {
		return "", err
	}
	defer twitterResponse.Body.Close()
	*twitterResponseCodeList = append(*twitterResponseCodeList, twitterResponse.Status)

	newTweetId, err := getTweetId(twitterResponse)
	if err != nil {
		return "", err
	}

	return newTweetId, nil
}

func sendReplyTweet(tweet string, twitterClient *http.Client, twitterResponseCodeList *[]string, previousTweetId string) (string, error) {
	tweetRequestBody, err := generateReplyRequestBody(tweet, previousTweetId)
	if err != nil {
		return "", err
	}

	twitterResponse, err := twitterClient.Post(url, "application/json", tweetRequestBody)
	if err != nil {
		return "", err
	}
	defer twitterResponse.Body.Close()
	*twitterResponseCodeList = append(*twitterResponseCodeList, twitterResponse.Status)

	newTweetId, err := getTweetId(twitterResponse)
	if err != nil {
		return "", err
	}

	return newTweetId, nil
}

func getTweetId(post *http.Response) (string, error) {
	var twitterResponseBody model.TwitterResponseBody
	twitterResponseBodyJSON, _ := ioutil.ReadAll(post.Body)
	err := json.Unmarshal(twitterResponseBodyJSON, &twitterResponseBody)
	return twitterResponseBody.Data.Id, err
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
