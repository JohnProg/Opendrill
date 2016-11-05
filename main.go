package main

import (
	"gopkg.in/mgo.v2"
	"encoding/json"
	"net/http"
	"log"
	"os"
	"./app/models"
	"./router"
)

var settings struct {
    Debug 			bool 	`json:"Debug"`
    Url  			string 	`json:"Url"`
    Port  			string 	`json:"Port"`
    DatabaseName  	string 	`json:"DatabaseName"`
    DatabaseUrl  	string 	`json:"DatabaseUrl"`
    Https  			bool 	`json:"Https"`
}

var (
	session *mgo.Session
	db      *mgo.Database
)

func init() {
	var err error

	configFile, err := os.Open("settings.json")
    if err != nil {
    	log.Fatalf("opening settings file", err.Error())
    }

    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&settings); err != nil {
    	log.Fatalf("parsing config file", err.Error())
    }

	session, err = mgo.Dial(settings.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", settings.DatabaseUrl)
	}
	session.SetMode(mgo.Strong, true)
	db = session.DB(settings.DatabaseName)
	models.SetDB(db)
}

func CloseSession() {
	session.Close()
}

func main() {

	defer session.Close()

	router.Init()
	log.Println("Running:", settings.Url)
	addr := settings.Url
	err := http.ListenAndServe(addr, nil)
	log.Println(err.Error())
}
