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
	api := twitter.InitTwitterApi()
	baseBlock := twitter.RetrieveOwnTweets("Sw_Saturn", api)
	result := markov.GenerateTweet(baseBlock)
	twitter.PostTweet(result, api)
}
