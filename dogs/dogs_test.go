package dogs

import (
	"fmt"
	"testing"
)

func ExampleYears1(t *testing.T) {
	fmt.Println(Years1(10))
	//Output:
	//100
}

func BenchmarkYears1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years1(10)
	}
}
func BenchmarkYears2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years2(10)
	}
}

func TestYears1(t *testing.T) {
	n := Years1(10)
	if n != 100 {
		t.Error("expected : 100, got : ", t)
	}
}
