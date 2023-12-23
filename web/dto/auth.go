package dto

type RegisterRequest struct {
	Email string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}
