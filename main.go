package main

import (
	"github.com/joho/godotenv"
	"log"
	"m/pkg/markov"
	"m/pkg/twitter"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	api := twitter.InitTwitterAPI()
	baseBlock := twitter.RetrieveOwnTweets("Sw_Saturn", api)
	result := markov.GenerateTweet(baseBlock)
	twitter.PostTweet(result, api)
	haiku := markov.GenerateHaiku(baseBlock)
	twitter.PostTweet(haiku, api)
}
