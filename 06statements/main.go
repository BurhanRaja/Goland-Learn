package main

import "fmt"

func main() {
	i := 2
	j := 2

	// If else
	if i > j {
		fmt.Println("Hello")
	} else if i == j {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Bye")
	}

	// Switch
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	tut := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	tut(7)
}
