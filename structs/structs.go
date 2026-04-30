package main

import (
	"errors"
	"fmt"
	"time"
)

// struct is used for grouping data
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt )
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Utility function to create a new user struct making it a pointer
func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required to create a user")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate(MM/DD/YYYY): ")

	// ..... Do something awesome with that gathered data!

	u, err := newUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	u.outputUserDetails()
	u.clearUserName()
	u.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
