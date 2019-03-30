package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joaquinicolas/twitter-hashtags/go/tweets-downloader/twitter"
)

var (
	key         = flag.String("key", "", "the consumer key provided by twitter")
	secret      = flag.String("secret", "", "consumer secret provided by twitter")
	token       = flag.String("token", "", "access token")
	tokenSecret = flag.String("tokenSecret", "", "")
)

func main() {
	flag.Parse()
	fmt.Println(twitter.Hi("Rob!"))
	listener := make(chan string)

	anaconda.SetConsumerKey(*key)
	anaconda.SetConsumerSecret(*secret)
	api := anaconda.NewTwitterApi(*token, *tokenSecret)

	searchTweets(listener, api)

	select {
		case
	}
}

func searchTweets(out chan string, api *anaconda.TwitterApi) {
	go func() {
		searchResult, err := api.GetSearch("golang", nil)
		if err != nil {
			panic(err)
		}
		for _, t := range searchResult.Statuses {
			fmt.Println(t.Text)
		}

	}()
}
