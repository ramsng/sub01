package sub01

import (
	"fmt"
	"math"

	"github.com/ramsng/sub02"
)

func Depref(a float64, b float64) int {
	fmt.Printf(sub02.Desc)
	fmt.Printf("** We are in package sub01 - inputs : %d ; %d\n", int(a), int(b))
	c := int((math.Pow(a, b)) * (math.Pow(b, a)))
	return c
}
