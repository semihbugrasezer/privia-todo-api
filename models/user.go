package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"` // "user" veya "admin"
}
