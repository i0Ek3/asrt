package asrt

import (
	"testing"
)

func TestAsrt(t *testing.T) {

	t.Run("Equal()", func(t *testing.T) {
		got := Equal(t, 10, 1+9)
		want := true

		if got != want {
			t.Logf("got %v equal with want %v", got, want)
		}
	})

	t.Run("NotEqual()", func(t *testing.T) {
		got := NotEqual(t, "yes", "no")
		want := true

		if got != want {
			t.Logf("got %v not equal with want %v", got, want)
		}
	})

	t.Run("NotNil()", func(t *testing.T) {
		got := NotNil(t, nil)
		want := true

		if got != want {
			t.Logf("got %v should not be nil", got)
		}
	})
}

func BenchmarkEqual(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		if Equal(t, i, i) {
			//
		}
	}
}

func BenchmarkNotEqual(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		if NotEqual(t, i, i-1) {
			//
		}
	}
}

func BenchmarkNotNil(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		if NotNil(t, nil) {
			//
		}
	}
}
