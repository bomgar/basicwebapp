package controllers

type Controllers struct {
	AuthController *AuthController
}

func Setup() *Controllers {
	return &Controllers{
		AuthController: &AuthController{},
	}
}
