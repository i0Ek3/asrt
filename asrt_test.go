package asrt

import (
	"testing"
)

func TestAsrt(t *testing.T) {
	got := []interface{}{}
	want := []interface{}{}

	Asrt(t, got, want)
}

func BenchmarkAsrt(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		Asrt(t, i, i)
	}
}
