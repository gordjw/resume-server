package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gordjw/resume-server/models"
)

func getRoles(w http.ResponseWriter, r *http.Request) {
	// Example data, to be replaced with a database call
	roles := map[string]models.Role{
		"propel":     models.NewRole("propel", "Senior Product Manager", "Propel Design", "A description goes here.", "2021", "Current"),
		"ag":         models.NewRole("ag", "Product Manager", "Department of Agriculture, Forestry and Fisheries", "A description goes here.", "2022", "2023"),
		"health":     models.NewRole("health", "Product Designer", "Department of Health and Aged Care", "A description goes here.", "2022", "2022"),
		"employment": models.NewRole("employment", "Product Coach and Delivery Manager", "Department of Education, Skills and Employment", "A description goes here.", "2021", "2022"),
		"dta":        models.NewRole("dta", "Product Manager", "Digital Transformation Agency", "A description goes here.", "2017", "2021"),
		"treasury":   models.NewRole("treasury", "Tech Lead, Senior Full-stack Developer", "Department of the Treasury", "A description goes here.", "2013", "2017"),
		"finance":    models.NewRole("finance", "Full-stack Developer", "Department of Finance", "A description goes here.", "2009", "2013"),
	}

	bytes, err := json.Marshal(roles)
	if err != nil {
		fmt.Printf("Error marshalling roles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func getRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Example data, to be replaced with a database call
	roles := map[string]models.Role{
		"propel":     models.NewRole("propel", "Senior Product Manager", "Propel Design", "A description goes here.", "2021", "Current"),
		"ag":         models.NewRole("ag", "Product Manager", "Department of Agriculture, Forestry and Fisheries", "A description goes here.", "2022", "2023"),
		"health":     models.NewRole("health", "Product Designer", "Department of Health and Aged Care", "A description goes here.", "2022", "2022"),
		"employment": models.NewRole("employment", "Product Coach and Delivery Manager", "Department of Education, Skills and Employment", "A description goes here.", "2021", "2022"),
		"dta":        models.NewRole("dta", "Product Manager", "Digital Transformation Agency", "A description goes here.", "2017", "2021"),
		"treasury":   models.NewRole("treasury", "Tech Lead, Senior Full-stack Developer", "Department of the Treasury", "A description goes here.", "2013", "2017"),
		"finance":    models.NewRole("finance", "Full-stack Developer", "Department of Finance", "A description goes here.", "2009", "2013"),
	}

	role := roles[id]

	if (role == models.Role{}) {
		fmt.Printf("Error finding role: %v\n", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(role)
	if err != nil {
		fmt.Printf("Error marshalling role: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
