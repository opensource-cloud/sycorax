package core

import "reflect"

type (
	Reflection struct {
		object interface{}
	}
)

// NewReflection returns a new Reflection instance
func NewReflection(object interface{}) *Reflection {
	return &Reflection{
		object: object,
	}
}

// GetJsonPropertyName returns a string.
// e.g object.entries -> [key, _].find => key === propertyName -> `json "propertyNameInJson"`
func (r *Reflection) GetJsonPropertyName(propertyName string) string {
	value := reflect.ValueOf(r.object)
	field, _ := value.Type().FieldByName(propertyName)
	return field.Tag.Get("json")
}
