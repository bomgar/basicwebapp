package dto

type RegisterRequest struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}