package receive

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func receive[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func ReceiveAndValidate[T any](r *http.Request, validator *validator.Validate) (T, error) {
	v, err := receive[T](r)
	if err != nil {
		return v, fmt.Errorf("receive json: %w", err)
	}

	validationErrors := validator.Struct(v)
	if validationErrors != nil {
		return v, fmt.Errorf("validation failed: %w", validationErrors)
	}
	return v, nil

}
