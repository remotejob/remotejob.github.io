package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/remotejob/docker-goreverseproxy/dbhandler"
	mgo "gopkg.in/mgo.v2"
)

func init() {

	log.Println("Sitemap init user", os.Getenv("SECRET_USERNAME"))
	mongodbuser = os.Getenv("SECRET_USERNAME")
	mongodbpass = os.Getenv("SECRET_PASSWORD")
	themes = "job"
	locale = "fi_FI"

}

//CreateArticelePage Create JSON result if mtitle not included
//100 article will be get
func CreateArticelePage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	mtitle := vars["mtitle"]

	mongoDBDialInfo := &mgo.DialInfo{

		Addrs:     []string{"mymongo-controller"},
		Timeout:   60 * time.Second,
		Database:  "admin",
		Username:  mongodbuser,
		Password:  mongodbpass,
		Mechanism: "SCRAM-SHA-1",
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	if mtitle == "" {

		log.Println("no mtitle")

		articles := dbhandler.GetAllForStatic(*dbsession, "kaukotyo.eu")
		json.NewEncoder(w).Encode(articles)

	} else {

		article := dbhandler.GetOneArticle(*dbsession, mtitle)

		json.NewEncoder(w).Encode(article)

	}
}
