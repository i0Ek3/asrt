package asrt

import (
	"reflect"
	"testing"
)

// Asrt outputs error message when got not equal with want
func Asrt(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %v", got, want)
	}
}

// Equal checks if the give value is equal
func Equal(t *testing.T, got, want interface{}) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	return false
}

// NotEqual checks if the give value is not equal
func NotEqual(t *testing.T, got, want interface{}) bool {
	if !reflect.DeepEqual(got, want) {
		return false
	}
	return true
}
