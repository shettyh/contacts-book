package api

import (
	"fmt"

	"github.com/shettyh/contacts-book/pkg/config"
)

// Serve will start the HTTP server with the configured router
func Serve() {
	conf := config.GetInstance()

	router := NewRouter()
	router.Run(fmt.Sprintf(":%d", conf.Port))
}
