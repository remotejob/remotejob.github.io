package main

import (
	//	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/remotejob/gmailhandler/server/handlers"
	"github.com/rs/cors"
	//	"github.com/gorilla/handlers"

	"net/http"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/sendemail", handlers.SendEmailHandler)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8090", c.Handler(gorillaRoute))
}
