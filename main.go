package main

import (
	"github.com/AntonBozhinov/techpodslearn/email"
	"github.com/AntonBozhinov/techpodslearn/learning"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

func main() {
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	mongoDb := os.Getenv("MONGO_DBNAME")
	emailSender := os.Getenv("EMAIL")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	port := os.Getenv("PORT")

	r := gin.Default()

	session, err := mgo.Dial(mongoConnectionString)
	if err != nil {
		log.Fatalf("could not establish db session: %v", err)
	}
	err = session.Ping()
	if err != nil {
		log.Fatalf("could not ping db: %v", err)
	}
	db := session.DB(mongoDb)

	sender := email.NewSender(emailSender, emailPassword)

	s := learning.NewServer(r, db, sender)
	s.Modules()
	s.ServeStaticFiles()
	log.Fatal(http.ListenAndServe(":"+port, s))
}
