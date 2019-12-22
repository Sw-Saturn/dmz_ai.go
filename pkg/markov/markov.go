package markov

import (
	crypto "crypto/rand"
	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"github.com/seehuhn/mt19937"
	"github.com/shogo82148/go-mecab"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

//BEGIN & END is used to detect beginning of sentence.
const (
	BEGIN = "__BEGIN_SENTENCE__"
	END   = "__END_SENTENCE__"
)

//DivideText is remove URLs.
func DivideText(text string) string {
	rep := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-…]+`)
	text = rep.ReplaceAllString(text, "")
	re := strings.NewReplacer("。", "", "．", "", ",", "", ".", "", "@", "", "#", "")
	result := re.Replace(text)
	return result
}

func _makeMarkovBlocks(analyzedText []string) [][]string {
	var markovBlock [][]string
	if len(analyzedText) < 3 {
		return markovBlock
	}
	head := []string{BEGIN, analyzedText[0], analyzedText[1]}
	markovBlock = append(markovBlock, head)

	for i := 0; i < len(analyzedText)-2; i++ {
		r := []string{analyzedText[i], analyzedText[i+1], analyzedText[i+2]}
		markovBlock = append(markovBlock, r)
	}
	foot := []string{analyzedText[len(analyzedText)-2], analyzedText[len(analyzedText)-1], END}
	markovBlock = append(markovBlock, foot)
	return markovBlock
}

func _extractWord(text string) []string {
	var words []string
	m, err := mecab.New(map[string]string{"dicdir": os.Getenv("MECAB_DIC_PATH")})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Destroy()

	node, err := m.ParseToNode(text)
	if err != nil {
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
	seed, _ := crypto.Int(crypto.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())

	for i, word := range markov[rng.Intn(len(markov))] {
		if i != 0 {
			result = append(result, word)
		}
	}
	return result
}


func _makeHaikuChain(markov [][]string, result []string) []string {
	seed, _ := crypto.Int(crypto.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())

	for i, word := range markov[rng.Intn(len(markov)-1)] {
		if i != 0 {
			result = append(result, word)
		}
	}
	return result
}

func _generateSentence(markovTable [][]string) string {
	var sentences []string
	var block [][]string
	firstTriplet := _getTriplet(BEGIN, markovTable)
	sentences = _makeChain(firstTriplet, sentences)
	count := 0
	for sentences[len(sentences)-1] != END {
		block = _getTriplet(sentences[len(sentences)-1], markovTable)
		if len(block) <= 0 {
			break
		}
		sentences = _makeChain(block, sentences)
		count++
		if count > 200 {
			break
		}
	}
	sentence := _joinSentences(sentences)
	return sentence
}

func _joinSentences(sentences []string) string {
	var result string
	for _, s := range sentences {
		if s == END {
			continue
		}
		result += s
	}
	return result
}

//GenerateTweet is generate Sentence.
func GenerateTweet(block []string) string {
	var markovBlocks [][]string
	for _, s := range block {
		_data := _extractWord(s)
		elems := _makeMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}
	s := _generateSentence(markovBlocks)
	for len(s) > 200 {
		s = _generateSentence(markovBlocks)
	}
	return s
}


func _countOfSounds(haikuParts string) int{
	t := tokenizer.New()
	tokens := t.Tokenize(haikuParts)
	count := 0
	for _, token := range tokens {
		features := token.Features()
		if token.Class == tokenizer.DUMMY {
			continue
		}
		if len(features) >= 8 {
			count += utf8.RuneCountInString(features[7])
		}
	}
	return count
}


func _getHaikuTriplet(markov [][]string) [][]string {
	var result [][]string
	for _, s := range markov {
			result = append(result, s)
	}
	return result
}


func _generateHaiku(markovTable [][]string) string {
	var sentences []string
	var haiku []string
	var sounds int
	for i := 0; i<=3; i++ {
		sentences = make([]string, 5)
		if i == 1 {
			firstTriplet := _getHaikuTriplet(markovTable)
			sentences = _makeHaikuChain(firstTriplet, sentences)
			sounds = 5
		} else if i == 2{
			firstTriplet := _getTriplet(BEGIN, markovTable)
			sentences = _makeHaikuChain(firstTriplet, sentences)
			sounds = 7
		} else if i == 3{
			firstTriplet := _getTriplet(BEGIN, markovTable)
			sentences = _makeHaikuChain(firstTriplet, sentences)
			sounds = 5
		}
		parts := _joinSentences(sentences)
		count := 0
		for _countOfSounds(parts) != sounds {
			firstTriplet := _getTriplet(BEGIN, markovTable)
			sentences = _makeHaikuChain(firstTriplet, sentences)
			parts = _joinSentences(sentences)
			count++
			if count > 200 {
				if i == 2 {
					parts = "なんもわからん"
				} else {
					parts = "最上川"
				}
			}
		}
		haiku = append(haiku, parts)
	}
	sentence := "ここで一句: " + _joinSentences(haiku)
	return sentence
}


//GenerateHaiku is generate a Haiku.
func GenerateHaiku(block []string) string {
	var markovBlocks [][]string
	for _, s := range block {
		_data := _extractWord(s)
		elems := _makeMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}
	s := _generateHaiku(markovBlocks)
	for len(s) == 17 {
		s = _generateHaiku(markovBlocks)
	}
	return s
}