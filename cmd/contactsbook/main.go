package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Start shutdown hook
	go shutdownhook()

	log.Print("Starting the HTTP server...")
	api.Serve()
}

func shutdownhook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for the exit signal
	<-c

	log.Print("Exit signal received. Cleaning up the resources...")
	db.GetSession().Close()

	os.Exit(0)
}
