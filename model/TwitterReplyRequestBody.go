package model

type TwitterReplyRequestBody struct {
	Text  string `json:"text"`
	Reply struct {
		InReplyToTweetId string `json:"in_reply_to_tweet_id"`
	} `json:"reply"`
}
