package main

import "m/pkg/markov"

const text = "ところで最近はお父さんお母さんの仕事帰りにCD等受け取りを頼むお子さんが多いですが、疎いジャンルの名前は覚えられない人もいますどうにか娘さんとのた。正解は『どついたれ本舗です』"

func main() {
	div := markov.DivideText(text)

}