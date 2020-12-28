package main

import (
	"fmt"
)

func returnTrue() bool {
	fmt.Println("\nReturning True")
	return true
}

func main() {

	guess := 30 // -5, 30, 50, 70
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Printf("\nGuess = %v. The guess must be between 1 - 100 !\n", guess)
	} else {
		fmt.Println("Guess is correct !!")
	}

}
