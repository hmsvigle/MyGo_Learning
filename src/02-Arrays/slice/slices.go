package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array :%v  \nLength: %v \n", a, len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	fmt.Println("\nslice = array[exclusive:inclusive]")
	b := a[:] // all elements
	fmt.Println(b)

	c := a[3:] // 4th Element till end
	fmt.Println(c)

	d := a[:6] // first 6 elements
	fmt.Println(d)

	e := a[3:6] // 4th, 5th, 6th element
	fmt.Println(e)

	fmt.Println("\nCreate Array with make(), which takes length & capacity as argument: ")
	fmt.Println("1 arg => arg = length = capacity")
	arr := make([]int, 4)
	fmt.Println(arr)
	fmt.Printf("Length: %v, Capacity: %v\n", len(arr), cap(arr))
	fmt.Println("2 args => arg =length  & arg2 = capacity")
	arr1 := make([]int, 4, 100)
	fmt.Println(arr1)
	fmt.Printf("Length: %v, Capacity: %v\n", len(arr1), cap(arr1))

}
