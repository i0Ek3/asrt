package asrt

import (
	"testing"
)

func TestAsrt(t *testing.T) {
	got := []interface{}{}
	want := []interface{}{}

	//Asrt(t, got, want)
	if Equal(t, got, want) {
		t.Logf("got %q equal with want %q", got, want)
	}

	//if NotEqual(t, got, want) {
	//	t.Logf("got %q not equal with want %q", got, want)
	//}
}

func BenchmarkAsrt(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		//Asrt(t, i, i)
		if Equal(t, i, i) {
			//
		}
		if NotEqual(t, i, i) {
			//
		}
	}
}
