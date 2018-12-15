package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rexitorg/sample-update-seed/request"
	"github.com/rexitorg/sample-update-seed/seed"
)

// LoadSeeds はseedをロードする
func LoadSeeds(w http.ResponseWriter, r *http.Request) {
	params := &request.PostPubSubParams{}
	err := parseRequest(r, params)
	if err != nil {
		log.Fatalf("Failed LoadSeeds: %#v", err)
	}

	fileName := params.GetFileName()
	err = seed.Load(r.Context(), fileName)
	if err != nil {
		log.Fatalf("Failed LoadSeeds: %#v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func parseRequest(r *http.Request, p interface{}) error {
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return err
	}

	return nil
}
