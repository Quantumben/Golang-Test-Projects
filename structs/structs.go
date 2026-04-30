package main

import (
	"fmt"
	"example.com/structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate(MM/DD/YYYY): ")

	// ..... Do something awesome with that gathered data!

	u, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	u.OutputUserDetails()
	u.ClearUserName()
	u.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
