package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"m/pkg/markov"
	"m/pkg/str"
	"net/url"
	"os"
)

func InitTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func RetrieveOwnTweets(username string, api *anaconda.TwitterApi) []string {
	tweetSources := []string{"TweetDeck", "Tweetbot", "Twitter"}
	var result []string
	v := url.Values{}
	v.Set("count", "200")
	v.Set("screen_name", username)
	v.Set("include_rts", "false")
	v.Set("include_entities", "0")
	v.Set("exclude_replies", "1")

	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tweets {
		if str.Contains(tweetSources, tweet.Source) {
			s := markov.DivideText(tweet.FullText)
			result = append(result, s)
		}
	}
	return result
}

func PostTweet(text string, api *anaconda.TwitterApi) {
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tweet.FullText)
}
