package main

import "fmt"

// Get compiled file using `go tool compile main.go`

func main() {
	fmt.Println("Hello, World!")
	concatenate_string()
	variables()
}

func concatenate_string() {
	var username string = "wagslane"
	var password string = "20947382822"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+" : "+password)
}

func variables() {
	// Simple declaration of variables
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// Short Variable declaration
	// limit := 123
	// temperature := 0.5
	// isFunny := true
	// empty := "Hello"

	// Type Interference
	j := 1
	i := j
	fmt.Println(i)

	// Convert Type
	accountAge := 2.6
	accountAgeInt := int64(accountAge)

	fmt.Println(accountAgeInt)

	// Contants
	const number = 12 // any type can be used here
}
