package main

import "fmt"

func main() {
	// Arrays -- fixed length, name types
	// var fruitArr [2]string
	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr := [2]string{"Grapes", "Pears"}
	fruitSlice := []string{"Grapes", "Pears", "Peaches"}

	// non-inclusive slice of a slice
	fmt.Println(fruitSlice[1:2])
	fmt.Println(fruitArr, len(fruitSlice))
}
