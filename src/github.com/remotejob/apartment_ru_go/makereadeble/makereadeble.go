package makereadeble

import (
	"math/rand"
	"strings"
	"time"

	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/comutils/str"
)

func Makehuman(mtext string) (string, string, string) {

	var title string
	var contents string
	var mcontents string
	var q int
	var sentenses []string

	mtexttokens := strings.Fields(mtext)

	title = ""
	var tmpline string
	q = gen.Random(11, 50)

	for i, token := range mtexttokens {

		if i <= 10 {
			title = title + " " + token
		}

		if i > 10 && i < len(mtexttokens) {

			if i <= q {

				tmpline = tmpline + " " + token

				if i == q {
					rand.Seed(time.Now().UTC().UnixNano())
					q = gen.Random(i+1, i+50)

					if len(tmpline) > 70 {

						sentenses = append(sentenses, str.UpcaseInitial(tmpline)+".")
						tmpline = ""
					}

				}

			}

		}

	}

	contentsQuant := len(sentenses) / 2

	for i, sentense := range sentenses {
		if i <= contentsQuant {

			contents = contents + " " + sentense

		} else {

			mcontents = mcontents + " " + sentense

		}

	}

	return title, strings.TrimSpace(contents), strings.TrimSpace(mcontents)

}
