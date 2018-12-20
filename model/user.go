package model

type User struct {
	ExternalUser
	Phone    string `json:"phone"`
	PersonID string `json:"person_id"`
}

type ExternalUser struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
}
