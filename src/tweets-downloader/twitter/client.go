package twitter

import (
	"flag"
	"fmt"
)

func createClient() {

	var accessToken string
	var accessTokenSecret string
	var consumerKey string
	var consumerKeySecret string

	flag.StringVar(&accessToken, "access-token", "TOKEN", "put your access token in here")
	flag.StringVar(&accessTokenSecret, "access-token-secret", "TOKEN", "put your access token secret in here")
	flag.StringVar(&consumerKey, "consumer-key", "KEY", "put your consumer key in here")
	flag.StringVar(&consumerKeySecret, "consumer-key-secret", "KEY", "put your consumer key secret in here")

	flag.Parse()

	fmt.Println("test" + accessToken + accessTokenSecret + consumerKey + consumerKeySecret)
	//	twitterClient := anaconda.NewTwitterApiWithCredentials(*accessToken, *accessTokenSecret, *consumerKey, *consumerKeySecret)
}
