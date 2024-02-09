package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting server")
	port := 8080

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello world"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
