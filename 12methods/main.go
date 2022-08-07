package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Website")

	tempUser := User{"Burhan", "mails4burhan@go.com", 9812987982, "Go123"}

	fmt.Println(tempUser)

	fmt.Printf("User information: %+v\n", tempUser)

	tempUser.getPhone()
	fmt.Println()
	tempUser.NewEmail()

	fmt.Printf("User information: %+v\n", tempUser)

}

type User struct {
	Username string
	Email    string
	Phone    int
	Password string
}

func (u *User) getPhone() {
	fmt.Printf("User phone number is %+v", u.Phone)
}

func (u *User) NewEmail() {
	u.Email = "hellogo@go.dev"
	fmt.Printf("User new email is %+v", u.Email)
}
