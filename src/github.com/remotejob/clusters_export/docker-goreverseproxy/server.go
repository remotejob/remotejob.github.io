// * Standart GOLANG service a part of JSON creation it's create XML service as well.
// * Smallest docker image used (don't forget GOLANG best for docker).
// * Documentation in docker-goreverseproxy

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/remotejob/docker-goreverseproxy/handlers"
	"github.com/remotejob/kaukotyoeu/handlers/robots"
	// _ "github.com/remotejob/godocker/statik"
)

var mongodbuser string
var mongodbpass string

func init() {

	log.Println("Start init")
	mongodbuser = os.Getenv("SECRET_USERNAME")
	mongodbpass = os.Getenv("SECRET_PASSWORD")

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}

	if _, err := os.Stat("/usr/share/nginx"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Println("/usr/share/nginx not exit ")

	} else {

		log.Println("/usr/share/nginx exist delete ")
		os.RemoveAll("/usr/share/nginx")

	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/robots.txt", robots.Generate)
	r.HandleFunc("/sitemap.xml", handlers.CheckServeSitemap)
	r.HandleFunc("/job/{locale}/{themes}", handlers.CreateArticelePage)
	r.HandleFunc("/job/{locale}/{themes}/{mtitle}.html", handlers.CreateArticelePage)
	log.Fatal(http.ListenAndServe(":8080", r))

}
