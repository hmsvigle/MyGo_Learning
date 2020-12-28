package main

import (
	"fmt"
)

func main() {
	// By default assigned false. so no initialization required for boolea variable.
	var b bool
	fmt.Printf("%v, %T\n", b, b)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("value of n after comparision - %v, %T\n", n, n)
	fmt.Printf("value of m after comparision - %v, %T\n", m, m)
}
