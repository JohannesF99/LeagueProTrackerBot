package config

import "github.com/magiconair/properties"

type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

func GetTwitterProperties() Credentials {
	props := loadProperties()
	return Credentials{
		first(props.Get("TWITTER_API_CONSUMER_KEY")),
		first(props.Get("TWITTER_API_CONSUMER_SECRET")),
		first(props.Get("TWITTER_API_TOKEN")),
		first(props.Get("TWITTER_API_TOKEN_SECRET")),
	}
}

func GetRiotKey() string {
	props := loadProperties()
	return first(props.Get("RIOT_API_KEY"))
}

func loadProperties() *properties.Properties {
	props, _ := properties.LoadFile("./config/api_details.properties", properties.UTF8)
	return props
}

func first(value string, _ bool) string {
	return value
}
