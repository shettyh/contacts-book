package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shettyh/contacts-book/pkg/api"
	"github.com/shettyh/contacts-book/pkg/config"
	"github.com/shettyh/contacts-book/pkg/db"
)

func main() {
	log.Print("Starting the Contacts book service...")

	log.Print("Initialize the configurations...")
	config.GetInstance()

	log.Print("Initializing the database...")
	db.GetSession()

	// Start shutdown hook
	go shutdownHook()

	log.Print("Starting the HTTP server...")
	api.Serve()
}

func shutdownHook() {
	exitSignalChan := make(chan os.Signal, 1)
	signal.Notify(exitSignalChan, os.Interrupt, syscall.SIGTERM)

	// Wait for the exit signal
	<-exitSignalChan

	log.Print("Exit signal received. Cleaning up the resources...")
	db.GetSession().Close()

	os.Exit(0)
}
