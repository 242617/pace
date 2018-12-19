package model

type User struct {
	ExternalUser
	Phone  string `json:"phone"`
	FaceID string `json:"face_id"`
}

type ExternalUser struct {
	Name string `json:"name"`
}
