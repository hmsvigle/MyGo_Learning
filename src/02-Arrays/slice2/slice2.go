package main

import (
	"fmt"
)

func main() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a), cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a), cap(a))
	fmt.Println("Copies old array & allocates a new memory; so capacity increases.")
	//fmt.Printf("\n Make can be used to create slices.\n Takes 2 arg - len & cap \n")

	a = append(a, 2, 3)
	fmt.Println(a)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a), cap(a))

	fmt.Println("Append slices of Array with different syntax:-")
	a = append(a, []int{5, 6, 7}...)
	fmt.Println(a)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a), cap(a))

	fmt.Println("** Changes to the slice & effect main underying Array: ")
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)

}
