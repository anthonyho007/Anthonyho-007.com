package main

import (
	"log"
	"os"

	"github.com/anthonyho007/GO-WebApp/middleware"
)

func main() {

	// acquire port
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT value must be define")
	}

	// start Web App
	middleware.App(port)
}
