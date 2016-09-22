package dbhandler

import (
	//	"fmt"
	"log"

	"github.com/remotejob/apartment_ru_go/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetAllForStatic coments
func GetAllForStatic(session mgo.Session) []domains.Articlefull {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")
	var results []domains.Articlefull
	err := c.Find(nil).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results

}

// GetOneArticle
func GetOneArticle(session mgo.Session, stitle string) domains.Article {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")

	var result domains.Article

	err := c.Find(bson.M{"stitle": stitle}).Select(bson.M{"created": 0, "updated": 0, "stitle": 0, "site": 0}).One(&result)
	if err != nil {

		log.Fatal(err)
		//		return
	}

	return result

}

func GetAllSitemaplinks(session mgo.Session) []domains.Sitemap_from_db {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")
	var results []domains.Sitemap_from_db
	err := c.Find(nil).Select(bson.M{"stitle": 1, "site": 1, "updated": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}

func GetAllUseful(session mgo.Session, themes string, locale string) []domains.Gphrase {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gkeywords").C("keywords")

	var results []domains.Gphrase

	err := c.Find(bson.M{"Themes": themes, "Locale": locale}).Select(bson.M{"Phrase": 1, "Rating": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}

// func InsetArticle(session mgo.Session, article entryHandler.Article) {
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("blog").C("articles")

// 	err := c.Insert(article)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
