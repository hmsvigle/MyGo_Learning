package main

import (
	"fmt"
)

func main() {
	switch 3 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Not one OR two")
	}

}
