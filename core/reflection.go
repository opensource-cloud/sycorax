package core

import "reflect"

type (
	Reflection struct {
	}
)

// NewReflection returns a new Reflection instance
func NewReflection() *Reflection {
	return &Reflection{}
}

// GetJsonPropertyName returns a string.
// e.g object.entries -> [key, _].find => $key === $propertyName -> `json "propertyNameInJson"`
func (r *Reflection) GetJsonPropertyName(structobject interface{}, propertyName string) string {
	value := reflect.ValueOf(structobject)
	field, _ := value.Type().FieldByName(propertyName)
	return field.Tag.Get("json")
}
