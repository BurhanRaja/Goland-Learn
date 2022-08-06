package main

import "fmt"

func main() {
	myNumber := 52

	var num *int = &myNumber
	fmt.Println("MyNumber Pointer value is: ", *num)

	ptr := &myNumber
	*ptr += 3
	fmt.Println("MyNumber Pointer value changed: ", myNumber)
}
