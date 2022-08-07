package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Website")

	tempUser := User{"Burhan", "mails4burhan@go.com", 9812987982, "Go123"}

	fmt.Println(tempUser)

	fmt.Printf("User information: %+v", tempUser)
}

type User struct {
	Username string
	Email    string
	Phone    int
	Password string
}
