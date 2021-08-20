package asrt

import (
	"reflect"
	"testing"
)

// deepequal wraps DeepEqual
func deepequal(t *testing.T, got, want interface{}) bool {
	if t && reflect.DeepEqual(got, want) {
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
	if t && value != nil {
		return true
	}
	return false
}

// Neg checks if prevalue equals given value neg
func Neg(t *testing.T, prepare, neg interface{}) bool {
    return deepequal(t, prepare, neg)
}

func AssertType(v interface{}) {
    switch typ := v.(type) {
    case int:
        fmt.Println(typ, "is int")
    case string:
        fmt.Println(typ, "is string")
    case nil:
        fmt.Println(typ, "is nil")
    default:
        fmt.Println(typ, "no type matched")
    }
}
