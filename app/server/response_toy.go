package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gordjw/resume-server/models"
)

func getToys(w http.ResponseWriter, r *http.Request) {
	// Example data, to be replaced with a database call
	toys := map[string]models.Toy{
		"fib":       models.NewToy("fib", "Fibonacci sequence calculator", "Algorithm", "some data"),
		"starfield": models.NewToy("starfield", "Starfield", "Javascript", "some data"),
	}

	bytes, err := json.Marshal(toys)
	if err != nil {
		fmt.Printf("Error marshalling toys: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func getToy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Example data, to be replaced with a database call
	toys := map[string]models.Toy{
		"fib":       models.NewToy("fib", "Fibonacci sequence calculator", "Algorithm", "some data"),
		"starfield": models.NewToy("starfield", "Starfield", "Javascript", "some data"),
	}

	toy := toys[id]

	if (toy == models.Toy{}) {
		fmt.Printf("Error finding toy: %v\n", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(toy)
	if err != nil {
		fmt.Printf("Error marshalling toy: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
