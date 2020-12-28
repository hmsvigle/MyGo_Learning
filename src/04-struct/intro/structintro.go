package main

import "fmt"

// Doctor is a structure
type Doctor struct {
	number    int
	actorName string
	companion []string
}

func main() {

	fmt.Println("1.Struct is a map with slices in it")

	aDoctor := Doctor{
		number:    3,
		actorName: "John Pet",
		companion: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName, aDoctor.companion[1])
}
