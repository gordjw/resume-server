package models

type Role struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Client      string `json:"client"`
	Description string `json:"description"`
	StartYear   string `json:"startYear"`
	EndYear     string `json:"EndYear"`
}

func NewRole(id, title, client, description, startYear, endYear string) Role {
	return Role{
		Id:          id,
		Title:       title,
		Client:      client,
		Description: description,
		StartYear:   startYear,
		EndYear:     endYear,
	}
}
