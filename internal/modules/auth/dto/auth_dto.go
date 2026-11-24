package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min:2"`
	Eamil    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min:6"`
}

type LoginRequesr struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" bindind:"required,min=6"`
}
