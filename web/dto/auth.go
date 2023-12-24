package dto

type RegisterRequest struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type LoginResponse struct {
	UserId int32 `json:"userId"`
}

type WhoAmIResponse struct {
	UserId int32 `json:"userId"`
}
