package server

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opensource-cloud/sycorax/core"
)

type (
	FieldError struct {
		Message string `json:"message"`
		Field   string `json:"field"`
		Value   string `json:"value"`
	}
	HttpError struct {
		Message string        `json:"message"`
		Code    string        `json:"code"`
		Cause   string        `json:"cause"`
		Fields  []*FieldError `json:"fields"`
	}
)

func validationErrorToText(e validator.FieldError) string {
	tag := e.Tag()
	field := e.Field()
	param := e.Param()
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", field, param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", field, param)
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", field, param)
	}
	return fmt.Sprintf("%s is not valid", field)
}

func (h *HttpError) ParseErrorsToFields(c *gin.Context, body interface{}) {
	r := core.NewReflection(body)
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			switch e.Type {
			case gin.ErrorTypeBind:
				errs := e.Err.(validator.ValidationErrors)
				for _, err := range errs {
					message := validationErrorToText(err)
					field := r.GetJsonPropertyName(err.Field())
					value, _ := json.Marshal(err.Value())
					h.Fields = append(h.Fields, &FieldError{
						Message: message,
						Field:   field,
						Value:   string(value),
					})
				}
			}
		}
	}
}

// NewHttpError returns a new instance of HttpError struct
func NewHttpError(m string, c string, e error) *HttpError {
	return &HttpError{
		Message: m,
		Code:    c,
		Cause:   e.Error(),
		Fields:  []*FieldError{},
	}
}

func NewInvalidSchemaError(e error) *HttpError {
	return NewHttpError("Invalid schema", "SCHEMA", e)
}

func NewInternalServerError(e error) *HttpError {
	return NewHttpError("Internal server error", "INTERNAL", e)
}
