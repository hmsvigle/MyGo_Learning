package main

import "fmt"

// Animal is a structure
type Animal struct {
	Name   string
	Origin string
}

// Bird is a structure
type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println("Embeded Structure:\n Its a Composition, not inheritence")
	fmt.Println(b)
	fmt.Println(b.Name)

}
