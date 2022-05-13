package interface_adapter

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

func ValidRequest(req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("failed to validate request object in controller: %w", err)
	}
	return nil
}
