package sentenses_tokenizer

import (
	"math/rand"
	"strings"
	"time"

	"github.com/shogo82148/go-shuffle"
	"gopkg.in/neurosnap/sentences.v1/english"
)

func Do_tokens(bfile []byte, quant int) []string {

	var retarray []string

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	sentences := tokenizer.Tokenize(string(bfile))
	//
	//	for _, s := range sentences {
	//		fmt.Println(s.Text)
	//	}

	var numberstoshuffle []int

	for num, _ := range sentences {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

	for i := 0; i < quant; i++ {
		//		fmt.Println(strings.TrimSpace(sentences[numberstoshuffle[i]].Text))
		retarray = append(retarray, strings.Replace(strings.TrimSpace(sentences[numberstoshuffle[i]].Text), "\n", "", -1))

	}

	return retarray
}
