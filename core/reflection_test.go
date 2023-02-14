package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/we-are-uranus/sycorax/core"
	"reflect"
	"testing"
)

type (
	myTestStruct struct {
		MyProp string
	}
)

func TestNewReflection(t *testing.T) {
	reflection := core.NewReflection(myTestStruct{})

	want := "Reflection"
	target := reflect.TypeOf(&reflection).Name()

	assert.Equal(t, want, target)
}
