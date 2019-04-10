package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	details = regexp.MustCompile(`(?:Landed|Boarding) ?(.*), ([A-Z]{3})`)
	apiKey = os.Getenv("TWITTER_API_KEY")
	apiSecret = os.Getenv("TWITTER_API_SECRET")
)

const (
	handle = "kevlinhenney"
	count = 500
)

func main() {
	config := &clientcredentials.Config{
		ClientID: apiKey,
		ClientSecret: apiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	params := &twitter.UserTimelineParams{
		ScreenName:      handle,
		ExcludeReplies:  twitter.Bool(true),
		IncludeRetweets: twitter.Bool(false),
		Count:           count,
	}

	httpClient := config.Client(context.Background())
	client := twitter.NewClient(httpClient)
	tweets, _, err := client.Timelines.UserTimeline(params)

	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		matches := details.FindStringSubmatch(tweet.Text)

		if len(matches) == 3 {
			process(matches[2], matches[1])
			break
		}
	}
}

func process(code, notes string) {
	if notes == "" {
		notes = "none"
	}

	fmt.Printf("Kevlin was last at: %s, additional notes: %s\n", code, notes)
}
