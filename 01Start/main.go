package main

import "fmt"

// Get compiled file using `go tool compile main.go`

func main() {
	fmt.Println("Hello, World!")
	concatenate_string()
}

func concatenate_string() {
	var username string = "wagslane"
	var password string = "20947382822"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+" : "+password)
}
