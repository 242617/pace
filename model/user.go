package model

type User struct {
	ExternalUser
	Name string `json:"name"`
}

type ExternalUser struct {
	Phone string `json:"phone"`
}
