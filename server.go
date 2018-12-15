package main

import (
	"log"
	"net/http"

	"github.com/rexitorg/sample-update-seed/router"
)

func main() {
	r := router.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed http listen and serve: %#v", err)
	}
}
