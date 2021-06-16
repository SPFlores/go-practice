package main

import "fmt"

func greeting(name string) string {
	return "Hello there " + name
}

// func getSum(num1 int, num2 int) int {
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Sammi"))
	fmt.Println(getSum(3, 4))
}
