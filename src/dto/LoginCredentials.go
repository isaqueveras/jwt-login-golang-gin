package dto

// Login credential
type LoginCredentials struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
}
