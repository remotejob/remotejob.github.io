package bookgen

import (
	"gopkg.in/mgo.v2"

	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	addrs := []string{"104.131.38.162"}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  "admin",
		Username:  "admin",
		Password:  "admin1Rel",
		Mechanism: "SCRAM-SHA-1",
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	Create(*dbsession, "realestate", "ru_RU", "/tmp/blog.txt")

}
