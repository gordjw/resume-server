package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer(port int) {
	apiRouter := chi.NewRouter()
	apiRouter.Get("/toys", getToys)
	apiRouter.Get("/toys/{id}", getToy)
	apiRouter.Get("/roles", getRoles)
	apiRouter.Get("/roles/{id}", getRole)

	r := chi.NewRouter()
	r.Use(middlewareCors)
	r.Use(middlewareLogger)

	r.Handle("/", http.FileServer(http.Dir("../client/build")))

	r.Mount("/api", apiRouter)

	fmt.Printf("Listening on %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
