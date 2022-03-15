package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("T", 3)
	expected := "TTT"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeat := Repeat("roma", 10)
	fmt.Println(repeat)
	// Output: romaromaromaromaromaromaromaromaromaroma
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("r", 0)
	}
}
