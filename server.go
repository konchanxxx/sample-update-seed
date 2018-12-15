package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rexitorg/sample-update-seed/router"
)

func main() {
	r := router.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Set default port: %s", port)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("Failed http listen and serve: %#v", err)
	}
}
