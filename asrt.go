package asrt

import (
	"reflect"
	"testing"
)

// deepequal wraps DeepEqual
func deepequal(t *testing.T, got, want interface{}) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	return false
}

// Asrt outputs error message when got not equal with want
func Asrt(t *testing.T, got, want interface{}) {
	if deepequal(t, got, want) {
		t.Errorf("got %q want %v", got, want)
	}
}

// Equal checks if the give value is equal
func Equal(t *testing.T, got, want interface{}) bool {
	if deepequal(t, got, want) {
		return true
	}
	return false
}

// NotEqual checks if the given value is not equal
func NotEqual(t *testing.T, got, want interface{}) bool {
	if !deepequal(t, got, want) {
		return false
	}
	return true
}

// NotNil checks if the given value is nil
func NotNil(t *testing.T, value interface{}) bool {
	if value != nil {
		return true
	}
	return false
}
