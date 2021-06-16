package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method for Person (value receiver)
func (hotdog Person) greet() string {
	// cannot have mixed data types in this, so have to convert other types to string
	return "Hello, my name is " + hotdog.firstName + " " + hotdog.lastName + ". I am " + strconv.Itoa(hotdog.age)
}

// hasBirthday method (pointer receiver), will be changing something
func (hotdog *Person) hasBirthday() {
	hotdog.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// a struct is sort of like a class in JS -- can have info or methods on it
	// init person using struct -- both are viable
	person1 := Person{firstName: "Samantha", lastName: "Flores", city: "Boston", gender: "F", age: 28}
	person2 := Person{"Bob","Johnson", "Boston", "M", 38}

	fmt.Println(person1, person2)

	// get single field
	// fmt.Println(person1.firstName)
	// edit field
	// person1.age++

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("williams")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
