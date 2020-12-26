package main

import (
	"fmt"
)

func main() {
	// int of 2 diff types not allowed to perform any airthmatic operations
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println("Sum: ", a+b)
	fmt.Println("Devide: ", a/b)
	fmt.Println("Reminder: ", a%b)

	fmt.Println("And operator: ", a&b) // 0010
	fmt.Println("OR operator", a|b)    // 1011
	fmt.Println("Exclusive OR: ", a^b) // 1001
	fmt.Println("And OR: ", a&^b)      // 0100
	// Bit Shifter

	fmt.Println("----Bit Shifters----")
	c := 8                             // 2^3
	fmt.Println("Left Shift: ", c<<3)  // 2^3  *  2^3 = 2^6
	fmt.Println("Right Shift: ", c>>3) // 2^3  /  2^3 = 2^0

	fmt.Println("---- Floating variables ----")
	// int of 2 diff types not allowed to perform any airthmatic operations
	// f is declared as float32. So we cant assign float64 no. It will throw error.
	f := 3.14
	f = 13.7e72
	f = 2.1e14

	fmt.Printf("%v, %T\n", f, f)

	fmt.Println("---- Complex Variables ----")

	var comp complex64 = 1 + 2i
	var comp2 complex64 = complex(2, 3)
	fmt.Printf("%v, %T\n", comp, comp)
	fmt.Printf("%v, %T\n", comp2, comp2)
}
