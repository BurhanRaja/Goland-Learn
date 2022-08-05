package main

import "fmt"

func main() {
	variables()
}

func variables() {
	// Simple declaration of variables
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// Short Variable declaration
	limit := 123
	temperature := 0.5
	isFunny := true
	empty := "Hello"

	fmt.Print(limit, temperature, isFunny, empty)

	// Type Interference
	j := 1
	i := j
	fmt.Println(i)

	// Convert Type
	accountAge := 2.6
	accountAgeInt := int64(accountAge)

	fmt.Println(accountAgeInt)

	// Contants
	// const number = 12 // any type can be used here
}
