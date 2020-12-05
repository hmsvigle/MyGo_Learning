// case3: Global & Local scoped variable

package main

import (
	"fmt"
)

var (
	actor     string = "Elizabeth Keen"
	companion string = "Sarah"
	series    string = "Blacklist"
	season    int    = 7
)

func main() {
	// Access global variable
	fmt.Println("----Access Global variable block----")
	fmt.Printf(" actor = %v\n companion = %v\n series=%v\n season = %v\n", actor, companion, series, season)
	// Reassign global variable
	season = 10
	fmt.Printf("\n Reassigned global variable i.e season = %v\n", season)
}
