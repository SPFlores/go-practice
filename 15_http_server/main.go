package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/hello-world", helloWorld)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
