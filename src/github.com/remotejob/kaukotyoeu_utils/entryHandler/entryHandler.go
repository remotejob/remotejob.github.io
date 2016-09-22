package entryHandler

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/gosimple/slug"
	"github.com/remotejob/apartment_ru_go/makereadeble"
	"github.com/remotejob/apartment_ru_go/mgenerator"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/comutils/str"

	"math/rand"

	"gopkg.in/mgo.v2"

	"log"
	"strings"
	"time"
)

type Article struct {
	Title     string
	Stitle    string
	Contents  string
	Mcontents string
	Site      string
	Author    string
	Created   time.Time
	Updated   time.Time
}

type Entryarticle struct {
	Modarticle Article
}

func NewEntryarticle() *Entryarticle {
	return &Entryarticle{Article{}}
}

func (article *Entryarticle) AddAuthor() {

	authorName := randomdata.FullName(randomdata.RandomGender)

	article.Modarticle.Author = authorName

}

func (article *Entryarticle) AddTitleStitleMcontents(bfile []byte, sites []string, uniqlinks map[string]struct{}) string {

	rand.Seed(time.Now().UTC().UnixNano())
	siteid := gen.Random(0, len(sites))

	mtext := mgenerator.Generate(bfile)

	title, contents, mcontents := makereadeble.Makehuman(mtext)

	stitle := slug.Make(title)
	article.Modarticle.Title = str.UpcaseInitial(title)
	article.Modarticle.Stitle = stitle
	article.Modarticle.Contents = contents
	article.Modarticle.Mcontents = mcontents
	article.Modarticle.Site = sites[siteid]

	return stitle

}

// func (article *Entryarticle) AddTags(tags []string) {

// 	var tagsquant = len(tags)
// 	var tags_str string = ""
// 	var tags_to_save []string
// 	for i := 0; i < 10; i++ {

// 		tagint := gen.Random(0, tagsquant)
// 		tags_to_save = append(tags_to_save, tags[tagint])
// 		tags_str = tags_str + " " + tags[tagint]

// 	}

// 	savetags.Saveinfile("createdtags.csv", tags_to_save)

// 	article.Modarticle.Tags = strings.TrimSpace(tags_str)

// }

func (article *Entryarticle) AddContents(sentenses []string) {

	var contents string = ""

	for _, sentens := range sentenses {

		contents = contents + " " + strings.Replace(sentens, "\n", "", -1)

	}

	article.Modarticle.Contents = str.UpcaseInitial(contents)

}

func (article *Entryarticle) InsertIntoDB(session mgo.Session) {

	backtime := gen.Random(0, 10000000)

	now := time.Now()

	then := now.Add(time.Duration(-backtime) * time.Second)
	articletodb := Article{article.Modarticle.Title, article.Modarticle.Stitle, article.Modarticle.Contents, article.Modarticle.Mcontents, article.Modarticle.Site, article.Modarticle.Author, then, then}
	//	dbhandler.InsetArticle(session, articletodb)
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")

	err := c.Insert(articletodb)
	if err != nil {
		log.Fatal(err)
	}

}
