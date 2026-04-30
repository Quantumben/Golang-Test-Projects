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

	u, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("test@example.com", "password123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

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
