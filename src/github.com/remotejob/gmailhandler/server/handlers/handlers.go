package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/remotejob/gmailhandler/domains"
	"github.com/remotejob/gmailhandler/sendgmail"
	"gopkg.in/gcfg.v1"
)

var glogin string
var gpass string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		glogin = cfg.Login.Glogin
		gpass = cfg.Pass.Gpass

	}

}

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {

	// urlParams := mux.Vars(r)
	// phone := urlParams["phone"]
	phone := r.Header.Get("phone")
	email := r.Header.Get("email")
	skype := r.Header.Get("skype")

	fmt.Println("phone", phone, "email", email, "skype", skype)

	sendgmail.Send(glogin, gpass, phone, email, skype)

	client := domains.Client{phone, email, skype}

	jData, err := json.Marshal(client)
	if err != nil {
		panic(err)

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}
