package main

import "fmt"

func main() {

	fmt.Println("\nDecleration: \n var <array_name> [<length>]<Type>")
	var a [3]string
	fmt.Printf("Empty Array : %v\n", a)

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("\nPrint each element separately: ")
	fmt.Println(a[0], a[1])
	fmt.Printf("\nPrint complete array: %v\n", a)

	fmt.Println("\nInitialization: \n <array_name> := [<length>]<Type>{ <enties> }")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("\nPrime Array: %v\n", primes)

	fmt.Println("Initialization without length: \n <array_name> := [<length>]<Type>{ <enties> }")
	grades := [...]int{67, 89, 88}
	fmt.Printf("\nPrime Array: %v\n", grades)

}
