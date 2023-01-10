package errors

import "fmt"

type (
	HttpError struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Error   string `json:"error"`
	}
)

// NewHttpError returns a new instance of HttpError struct
func NewHttpError(m string, c string, e string) *HttpError {
	return &HttpError{
		Message: m,
		Code:    c,
		Error:   e,
	}
}

func NewInvalidSchemaError(e error) *HttpError {
	return NewHttpError("Invalid schema", "SCHEMA", fmt.Sprintf(e.Error()))
}

func NewInternalServerError(e error) *HttpError {
	return NewHttpError("Internal server error", "INTERNAL", fmt.Sprintf(e.Error()))
}
