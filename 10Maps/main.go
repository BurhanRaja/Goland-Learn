package main

import "fmt"

func main() {
	languages := make(map[int]string)

	languages[1] = "Java"
	languages[2] = "Javascript"
	languages[3] = "Go"

	fmt.Println(languages)

	fmt.Println(languages[3])
}
