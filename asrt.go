package asrt

import (
    "fmt"
	"reflect"
	"testing"
)

// T ready for generic 
type (
    T  interface{}
)

// deepequal wraps DeepEqual
func deepequal(got, want T) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	return false
}

// Asrt outputs error message when got not equal with want
func Asrt(t *testing.T, got, want T) {
	if deepequal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Equal checks if the give value is equal
func Equal(got, want T) bool {
	if deepequal(got, want) {
		return true
	}
	return false
}

// NotEqual checks if the given value is not equal
func NotEqual(got, want T) bool {
	if !deepequal(got, want) {
		return false
	}
	return true
}

// NotNil checks if the given value is nil
func NotNil(value T) bool {
	if value != nil {
		return true
	}
	return false
}

// Neg checks if prevalue equals given value neg
func Neg(prepare, neg T) bool {
    return deepequal(prepare, neg)
}

// AssertType asserts the value type
func AssertType(t T) string {
    switch t := t.(type) {
    case nil:
        return "NULL"
    case int, uint:
        return fmt.Sprintf("%d", t)
    case bool:
        if t {
            return "TRUE"
        }
        return "FALSE"
    case string:
        return fmt.Sprintf("%s", t)
    default:
        panic(fmt.Sprintf("unexpected type %T: %v", t, t))
    }
}
