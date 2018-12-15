package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rexitorg/sample-update-seed/handler"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Post("/_ah/push-handlers/seeds", handler.LoadSeeds)

	return r
}
