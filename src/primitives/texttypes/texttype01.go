package main

import (
	"fmt"
)

func main() {

	s1 := "This is String1"
	s2 := "String2"
	b := []byte(s1)

	fmt.Printf("Concatenate strings: %v, %T\n", s1+s2, s1+s2)
	// when a file(consistes of words ie. bytes are passed to another program)
	fmt.Printf("String to byte: %v, %T\n", b, b)

	// Runes data type:-
	fmt.Println("---- Rune Data Type: Int32 data type----")
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
