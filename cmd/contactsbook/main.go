package main

import (
	"log"
	"os"

	"github.com/shettyh/contacts-book/pkg/config"

	"github.com/shettyh/contacts-book/pkg/db"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shettyh/contacts-book/pkg/api"
)

func main() {
	log.Print("Starting the Contacts book service...")

	log.Print("Initialize the configurations...")
	config.GetInstance()

	log.Print("Initializing the database...")
	db.GetSession()

	log.Print("Starting the HTTP server...")
	//TODO: remove this harcoded values
	os.Setenv("CB_DBHOST", "localhost")
	api.Serve()
}
