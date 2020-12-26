package main

import "fmt"

const a int = 27

// constant block defined
const (
	// iota = enumerated constant. iota assigns numerical sequence to constants (0 1 2 3 ...)
	x = iota
	y
	z = iota
)

func main() {

	const a int = 42

	fmt.Println("\n inner decleration can change the global constant value.")
	fmt.Printf("constant a = %v, %T\n", a, a)

	var b = 27                                            // var b int32 = 27
	fmt.Printf("change type result = %v, %T\n", a+b, a+b) // change type is not allowed

	fmt.Println("\niota is enumerated constant. iota assigns numerical sequence to constants (0 1 2 3 ...)")
	fmt.Printf("constant x = %v, y = %v, z = %v \n", x, y, z)
	// Constants cannot be declared using the := syntax.

}
