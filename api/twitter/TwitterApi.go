package twitter

import (
	"bytes"
	"encoding/json"
	"github.com/dghubble/oauth1"
	"net/http"
)

const url = "https://api.twitter.com/2/tweets"

type Credentials struct {
	consumerKey    string
	consumerSecret string
	token          string
	tokenSecret    string
}

func Tweet(message string) string {
	client := getClient(getCreds())
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

func getClient(creds Credentials) *http.Client {
	config := oauth1.NewConfig(creds.consumerKey, creds.consumerSecret)
	token := oauth1.NewToken(creds.token, creds.tokenSecret)
	return config.Client(oauth1.NoContext, token)
}

func getCreds() Credentials {
	return Credentials{
		"consumerKey",
		"consumerSecret",
		"token",
		"tokenSecret",
	}
}
