package asrt

import (
	"reflect"
	"testing"
)

// Asrt checks if the give value is equal
func Asrt(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
