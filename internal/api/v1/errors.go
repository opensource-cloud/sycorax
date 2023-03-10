package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opensource-cloud/sycorax/internal/core"
)

type (
	FieldError struct {
		Message string `json:"message"`
		Field   string `json:"field"`
		Value   string `json:"value"`
	}
	SycoraxError struct {
		Message     string        `json:"message"`
		TypeOfError string        `json:"type_of_error"`
		Reason      string        `json:"reason"`
		Fields      []*FieldError `json:"fields"`
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

func (object *SycoraxError) ParseErrorsToFields(c *gin.Context, body interface{}) {
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
					object.Fields = append(object.Fields, &FieldError{
						Message: message,
						Field:   field,
						Value:   string(value),
					})
				}
			}
		}
	}
}

func (object *SycoraxError) AddField(field *FieldError) *SycoraxError {
	object.Fields = append(object.Fields, field)
	return object
}

// NewSycoraxError returns a new instance of SycoraxError struct
func NewSycoraxError(m string, t string, e error) *SycoraxError {
	return &SycoraxError{
		Message:     m,
		TypeOfError: t,
		Reason:      e.Error(),
		Fields:      []*FieldError{},
	}
}

func NewInvalidSchemaError(e error) *SycoraxError {
	return NewSycoraxError("Invalid schema", "SCHEMA", e)
}

func NewUnprocessableEntity(e error) *SycoraxError {
	return NewSycoraxError("Unprocessable entity", "ENTITY", e)
}

func NewInternalServerError(e error) *SycoraxError {
	return NewSycoraxError("Internal Server Error", "INTERNAL", e)
}
