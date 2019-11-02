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
	var markovBlocks [][]string

	for _, s := range baseBlock {
		_data := markov.ExtractWord(s)
		elems := markov.MakeMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}

	elemsSet := markov.GenerateSentence(markovBlocks)
	twitter.PostTweet(elemsSet, api)
}
