package markov

import (
	crand "crypto/rand"
	"github.com/seehuhn/mt19937"
	"github.com/shogo82148/go-mecab"
	"log"
	"math"
	"math/big"
	"math/rand"
	"strings"
)

const (
	ipadic = "/usr/local/lib/mecab/dic/mecab-ipadic-neologd"
	BEGIN = "__BEGIN_SENTENCE__"
	END = "__END_SENTENCE__"
)

func DivideText(text string) string{
	rep := strings.NewReplacer("。","","．","",",","",".","")
	return rep.Replace(text)
}

func MakeMarkovBlocks(analyzedText []string) [][]string {
	var markovBlock [][]string
	if len(analyzedText) < 3{
		return markovBlock
	}
	head := []string{BEGIN,analyzedText[0],analyzedText[1]}
	markovBlock = append(markovBlock, head)

	for i := 0; i < len(analyzedText) - 2; i++ {
		r := []string{analyzedText[i],analyzedText[i+1],analyzedText[i+2]}
		markovBlock = append(markovBlock, r)
	}
	foot := []string{analyzedText[len(analyzedText)-2], analyzedText[len(analyzedText)-1], END}
	markovBlock = append(markovBlock, foot)
	return markovBlock
}

func ExtractWord(text string) []string {
	var words []string
	m, err := mecab.New(map[string]string{"dicdir": ipadic})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Destroy()

	node, err := m.ParseToNode(text)
	if err != nil{
		log.Fatal(err)
	}
	for ; !node.IsZero(); node = node.Next() {
		if node.Surface() != "" {
			words = append(words, node.Surface())
		}
	}
	return words
}

func _getTriplet(target string, markov [][]string) [][]string {
	var result [][]string
	for _, s := range markov {
		if s[0] == target {
			result = append(result, s)
		}
	}
	return result
}

func _makeChain(markov [][]string, result []string) []string {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())

	for i, word := range markov[rng.Intn(len(markov))]{
		if i != 0 {
			result = append(result, word)
		}
	}
	return result
}

func generateSentence(chain []string, markovTable [][]string) string{
	var sentence []string
	var block [][]string
	firstTriplet := _getTriplet(BEGIN, markovTable)
	sentence = _makeChain(firstTriplet, sentence)
	for sentence[len(sentence)-1] != END {
		block = _getTriplet(sentence[len(sentence)-1], markovTable)
		if len(block) <= 0 {
			break
		}
		sentence = _makeChain(firstTriplet, sentence)
	}
	return strings.Join(sentence, "")
}
