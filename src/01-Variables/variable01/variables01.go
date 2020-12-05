package main

import "fmt"

func main() {
	// case1: declare var & assign value later
	var i int
	i = 42
	fmt.Println(i)
	// case2: declare & assign a variable
	var j int = 27
	fmt.Printf("j = %v, %T\n", j, j)

}
