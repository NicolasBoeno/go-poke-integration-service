package errors

import "fmt"

const (
	ErrCodeNotFound      = "POKEMON_NOT_FOUND"
	ErrCodeInvalidInput  = "INVALID_INPUT"
	ErrCodeAPIError      = "API_ERROR"
	ErrCodeInternalError = "INTERNAL_ERROR"
)

func NewPokemonNotFoundError(name string) CustomError {
	return &ServiceError{
		code:    ErrCodeNotFound,
		message: fmt.Sprintf("pokemon not found: %s", name),
	}
}

func NewInvalidInputError(message string) CustomError {
	return &ServiceError{
		code:    ErrCodeInvalidInput,
		message: message,
	}
}

func NewAPIError(message string, err error) CustomError {
	return &ServiceError{
		code:    ErrCodeAPIError,
		message: message,
		err:     err,
	}
}

func NewInternalError(message string, err error) CustomError {
	return &ServiceError{
		code:    ErrCodeInternalError,
		message: message,
		err:     err,
	}
}
