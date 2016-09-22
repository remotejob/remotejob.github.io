package bookgen

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/remotejob/apartment_ru_go/dbhandler"
	"github.com/shogo82148/go-shuffle"
	"gopkg.in/mgo.v2"
)

func Create(session mgo.Session, themes string, locale string, filename string) {

	if _, err := os.Stat(filename); !os.IsNotExist(err) {

		err := os.Remove(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	articles := dbhandler.GetAllUseful(session, themes, locale)

	var numberstoshuffle []int

	for num := range articles {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

	for _, i := range numberstoshuffle {

		// moto := ""

		// for _, tag := range articles[i].Tags {

		// 	moto = moto + " " + tag

		// }

		paragraph := articles[i].Phrase + "\n"

		if _, err = f.WriteString(paragraph); err != nil {
			panic(err)
		}
	}
	f.Close()
}
