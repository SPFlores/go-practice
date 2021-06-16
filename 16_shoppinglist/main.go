package main

import (
	// "fmt"
	"net/http"

	"github.com/SPFlores/go_crash_course/16_shoppinglist/api"
)

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// w.Write([]byte("hello world"))
// }

func main() {
	// http.HandleFunc("/hello-world", helloWorld)
	// fmt.Println("Server is running...")
	// http.ListenAndServe(":8080", nil)
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
