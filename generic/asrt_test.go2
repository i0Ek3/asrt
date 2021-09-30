package asrt

import (
	"testing"
)

func TestAsrt(t *testing.T) {

	t.Run("Equal()", func(t *testing.T) {
		got := Equal(10, "1+9")
		want := true 

        Asrt(t, got, want)
	})

	t.Run("NotEqual()", func(t *testing.T) {
		got := NotEqual("yes", "no")
		want := true

        Asrt(t, got, want)
	})

	t.Run("NotNil()", func(t *testing.T) {
		got := NotNil(nil)
		want := true
        Asrt(t, got, want)
	})
}

func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Equal(i, i) {
			//
		}
	}
}

func BenchmarkNotEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if NotEqual(i, i-1) {
			//
		}
	}
}

func BenchmarkNotNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if NotNil(nil) {
			//
		}
	}
}
