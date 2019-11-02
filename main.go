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

func generateTweet(block []string) string {
	var markovBlocks [][]string
	for _, s := range block {
		_data := markov.ExtractWord(s)
		elems := markov.MakeMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}
	s := markov.GenerateSentence(markovBlocks)
	return s
}

func main() {
	loadEnv()
	api := twitter.InitTwitterApi()
	baseBlock := twitter.RetrieveOwnTweets("Sw_Saturn", api)
	result := generateTweet(baseBlock)
	for len(result) > 200 {
		result = generateTweet(baseBlock)
	}
	twitter.PostTweet(result, api)
}
