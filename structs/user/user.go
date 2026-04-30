package user

import (
	"errors"
	"fmt"
	"time"
)

// struct is used for grouping data
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct embedding is a way to include one struct within another struct, it allows us to reuse the fields and methods of the embedded struct in the outer struct without having to explicitly define them again. In this example, we have an Admin struct that embeds the User struct, which means that Admin has access to all the fields and methods of User without having to redefine them. This is a powerful feature of Go that promotes code reuse and composition.
// it's called inheritance in php
type Admin struct {
	email string
	password string
	User    User
} 


func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "User",
			birthDate: "01/01/1970",
			createdAt: time.Now(), 
		},
	}
}

// Utility function to create a new user struct making it a pointer
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required to create a user")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}


