package main

import "fmt"

func main() {
	// Pointers allow you to point to the memory address of a value
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)

	// Use * to read value from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// may use pointers if you need a lot of data passed. passing locations may be faster than passing full data
}
