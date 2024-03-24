package cm

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Fatal(err)
		}
	})
	return r
}

func Server() {
	err := http.ListenAndServe(":3000", Routes())
	if err != nil {
		log.Fatal(err)
	}
}
