package main

import "fmt"

func main() {

	fmt.Println("1.Map is 'key:value' pair\n2.output is not sequential.\n3.search keyfor value.")
	m := map[string]int{}
	fmt.Println(m)

	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 23451,
		"Texas":      26734,
		"Ohio":       3726477,
	}

	fmt.Println(statePopulation)
	fmt.Println(statePopulation["California"])
	fmt.Println(statePopulation["oho"])
	fmt.Println("if not sure, if key is available or not, use pop, ok:")
	pop, ok := statePopulation["Ohio"]
	fmt.Println(pop, ok)

	pop3, ok := statePopulation["Eho"]
	fmt.Println(pop3, ok)
}
