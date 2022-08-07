package main

import "fmt"

func main() {
	defer fmt.Println("World")
	myDefer()
	fmt.Println("Hello")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
