package models

// User is info
type User struct {
	Name  string `json:"name"`
	Token string `json:"json"`
}

// Secret is
type Secret struct {
	Key string `json:"key"`
	User
}
