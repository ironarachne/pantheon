package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/pantheon"
	"github.com/ironarachne/random"
)

func getPantheon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newPantheon pantheon.Pantheon

	random.SeedFromString(id)

	newPantheon = pantheon.GeneratePantheon(20)

	json.NewEncoder(w).Encode(newPantheon)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pantheongend"))
	})

	r.Get("/{id}", getPantheon)

	fmt.Println("Pantheon Generator API is online.")
	log.Fatal(http.ListenAndServe(":7535", r))
}
