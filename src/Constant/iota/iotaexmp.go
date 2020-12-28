package main

import "fmt"

const (
	m = 1 << iota
	n
	o
)

const (
	_ = iota
	// KB Kilo Bytes
	KB = 1 << (10 * iota)
	// MB Megabytes
	MB
	// GB GigaBytes
	GB
)

func main() {

	fmt.Println("Constants cannot be declared using the := syntax")

	fmt.Println("\nBitshift i.e '<<' ==>  2^")
	fmt.Println("m = 2^0, n = 2^1, o=2^2")
	fmt.Printf("m = %v\nn = %v\no = %v\n", m, n, o)

	fmt.Println("\n_= 2^0, KB = 2^10, MB = 2^100, GB=2^1000")
	fmt.Printf("KB = %v\nMB = %v\nGB = %v\n", KB, MB, GB)

	var filesize = 4000000000.0
	fmt.Printf("filesize = %.2f GB\n", filesize/GB)

}
