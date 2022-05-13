package interface_adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

type OKResponseBody struct {
	Messages []string `json:"messages"`
}

type ErrorResponseBody struct {
	Messages []string `json:"messages"`
}

func handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, model.ErrInternalServerError):
		return c.JSON(http.StatusInternalServerError, ErrorResponseBody{Messages: []string{model.ErrInternalServerError.Error()}})
	case errors.Is(err, model.ErrNotFound):
		return c.JSON(http.StatusNotFound, ErrorResponseBody{Messages: []string{model.ErrNotFound.Error()}})
	case errors.Is(err, model.ErrConflict):
		return c.JSON(http.StatusConflict, ErrorResponseBody{Messages: []string{model.ErrConflict.Error()}})
	case errors.As(err, &validator.ValidationErrors{}):
		// validator.ValidationErrorsに型アサーション可能になるまでUnwrapする
		unwrapedErr := func() error {
			for {
				if _, ok := err.(validator.ValidationErrors); ok {
					return err
				}
				err = errors.Unwrap(err)
			}
		}()

		type ValidationErr struct {
			Field        string      `json:"field"`
			InvalidValue interface{} `json:"invalid_value"`
			Tag          string      `json:"tag"`
			Param        string      `json:"param"`
		}
		var validationErrs []ValidationErr
		for _, e := range unwrapedErr.(validator.ValidationErrors) {
			validationErrs = append(
				validationErrs,
				ValidationErr{Field: e.Field(), InvalidValue: e.Value(), Tag: e.Tag(), Param: e.Param()},
			)
		}

		validErrRespByte, err := json.Marshal(validationErrs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponseBody{Messages: []string{model.ErrInternalServerError.Error()}})
		}

		return c.JSON(http.StatusBadRequest, ErrorResponseBody{Messages: []string{fmt.Sprintf("validation error has occured: %s", string(validErrRespByte))}})

	default:
		return c.JSON(http.StatusInternalServerError, ErrorResponseBody{Messages: []string{model.ErrInternalServerError.Error()}})
	}

}
