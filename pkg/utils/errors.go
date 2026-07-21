package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definicion de estructura de respuestas (errores de validacion)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Conversion de los errores binnding  a txt
func GetValidationError(err error) []ApiError {
	var msj validator.ValidationErrors
	if errors.As(err, &msj) {
		out := make([]ApiError, len(msj))
		for i, fe := range msj {
			out[i] = ApiError{
				Field:   fe.Field(),
				Message: msgForTag(fe),
			}
		}
		return out
	}
	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Este campo es obligatorio"
	case "max":
		return fmt.Sprintf("El valor excede el máximo permitido (%s)", fe.Param())
	case "min":
		return fmt.Sprintf("El valor es inferior a la longitud permitida (mínimo %s)", fe.Param())
	case "numeric":
		return "Este campo solo debe contener números"
	case "len":
		return fmt.Sprintf("Este campo debe tener exactamente %s caracteres", fe.Param())
	}
	return "Valor inválido"
}
