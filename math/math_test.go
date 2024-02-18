package math

import "testing"

func TestMax(t *testing.T) {
	//var minarr []float64;
	min := Max([]float64{1, 2, 3, 4, 5, 6})
	if min != 5 {
		t.Error("Expected 5, got :", min)
	}
}
