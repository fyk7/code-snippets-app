package interface_adapter

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

func isRequestValid(req interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, fmt.Errorf("failed to validate request object in controller: %w", err)
	}
	return true, nil
}
