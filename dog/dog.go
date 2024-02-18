package dog

import "fmt"

func Years(a *int) int {
	fmt.Println("Code in Dog Package")
	b := *a
	return b * 7
}
