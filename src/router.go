package cm

import (
	_ "github.com/campaign-manager/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Campaign Manager API
// @version 0.1
// @description Computation Campaign Manager REST API documentation.
// @BasePath /api/v1

// @contact.name Pascal Post
// @contact.url https://github.com/pascalPost/campaign-manager
// @contact.email pascal.post@mailbox.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Routes returns a chi router
func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/test", GetTest)

	r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})
	r.Get("/swagger*", httpSwagger.Handler())

	return r
}

func Server() {
	r := chi.NewRouter()
	r.Mount("/api/v1", Routes())
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}

// GetTest can be used to test the connection
// @Summary can be used to test the connection
// @Description GetTest returns OK
// @Router /test [get]
// @Success 200 {string}  string    "OK"
func GetTest(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err)
	}
}
