package main

import (
	"encoding/json"
	"fmt"
	"github.com/campaign-manager/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"os"
)

func main() {

	// TODO set this based on env variable
	addr := "localhost:3000"

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{

		// TODO set this based on env variable
		AllowedOrigins: []string{"http://localhost:5173"},

		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// TODO activate this
	//r.Use(nethttpmiddleware.OapiRequestValidator(swagger))

	// TODO read this from config
	strictServer := api.NewServer("./work")

	server := api.NewStrictHandler(strictServer, nil)

	handler := api.Handler(server)

	r.Mount("/", handler)

	//r.Get("/api/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/yaml")
	//	err := json.NewEncoder(w).Encode(oapi.Spec)
	//	if err != nil {
	//		if c.Dev() {
	//			http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//
	//		http.Error(w, "Internal error", http.StatusInternalServerError)
	//		return
	//	}
	//})
	//swh := swgui.NewHandler(filePathToSchema, urlToSchema, "/")

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/api/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(swagger)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	//hh := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(rw, r, "/api/openapi.yaml")
	//})

	//r.Handle("/", hh)

	//r.Mount("/api/docs", swgui.New(
	//	"test",
	//	"/api/openapi.json",
	//	"/api/docs/",
	//))
	//println("docs at http://" + addr + "/api/docs")

	//h := api.HandlerFromMux(handler, r)
	//
	//s := &http.Server{
	//	Handler: h,
	//	Addr:    addr,
	//}
	//
	//slog.Info("Server running at", "address", addr)
	//log.Fatal(s.ListenAndServe())

	http.ListenAndServe(addr, r)
}
