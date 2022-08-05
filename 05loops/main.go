package main

import "fmt"

func main() {
	// Kind of while loop
	i := 1
	for i < 6 {
		fmt.Printf("Hello %v\n", i)
		i += 1
	}

	// For loop
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	// For loop for string
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}
	for h, j := range rvariable {
		fmt.Println(h, j)
	}
}
