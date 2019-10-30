package main

import (
	"fmt"
	"m/pkg/markov"
)

const text = "ところで最近はお父さんお母さんの仕事帰りにCD等受け取りを頼むお子さんが多いですが、疎いジャンルの名前は覚えられない人もいますどうにか娘さんとのた。正解は『どついたれ本舗です』"
const text1 = "メロスは激怒した。必ず、かの邪智暴虐じゃちぼうぎゃくの王を除かなければならぬと決意した。メロスには政治がわからぬ。メロスは、村の牧人である。"
func main() {
	var baseBlock []string
	baseBlock = append(baseBlock, markov.DivideText(text))
	baseBlock = append(baseBlock, markov.DivideText(text1))
	var markovBlocks [][]string

	for _, s := range baseBlock {
		_data := markov.ExtractWord(s)
		elems := markov.MakeMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}

	elemsSet := markov.GenerateSentence(markovBlocks)
	fmt.Println(elemsSet)
}