package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The Test is true.")
	}
	// Example 2: with initializer & its syntax.
	fmt.Println("\nSyntanx:\n if <Initializer> ; <boolean condition> {\n \t<action> }")
	fmt.Println("condition has to be boolean in above syntax. ")
	if number := 50; true {
		fmt.Println(number)
	}
	// Example 3:
	number := 50
	guess := -5 // -5, 30, 50, 70
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 - 100 !")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("You Got it !!")
		}

		fmt.Println(number <= guess, number >= guess, number != guess)
	}

}
