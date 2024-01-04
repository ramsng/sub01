package sub01

import (
	"fmt"
	"math"
)

func Depref(a float64, b float64) int {
	fmt.Printf("** We are in package sub01 - inputs : %v ; %v", int(a), int(b))
	c := int((math.Pow(a, b)) * (math.Pow(b, a)))
	return c
}
