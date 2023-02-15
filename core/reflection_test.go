package core_test

import (
	"reflect"
	"testing"

	"github.com/opensource-cloud/sycorax/core"
	"github.com/stretchr/testify/assert"
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
