package main

import "fmt"

func main() {
	// var name string = "Sam"
	// can assign 2 variables on the same line by separating with comas using := syntax
	name := "Sam"
	var age int = 28
	const isCool bool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
}
