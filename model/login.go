package models

type LoginInput struct {
	Email     string `json:"email" binding:"omitempty,email"`
    Handphone string `json:"handphone" binding:"omitempty"`
    Password  string `json:"password" binding:"required"`
}
