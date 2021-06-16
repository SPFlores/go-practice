package main

import "fmt"

func main() {
	// maps are key-value pairs
	// emails := make(map[string]string)

	// assign key-values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// declare map and add key-values at same time
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com", "Mike":"mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// delete something from a map
	delete(emails, "Bob")
	fmt.Println(emails)
}
