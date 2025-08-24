package models

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}