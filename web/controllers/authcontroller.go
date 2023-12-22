package controllers

import "net/http"

type AuthController struct {
}

func (c *AuthController) WhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("don't know"))
}
