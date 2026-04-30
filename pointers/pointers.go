package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int 
	agePointer = &age //pointer variable that holds the address of age variable

	fmt.Println("Age", *agePointer)
	// fmt.Println(&age)
}

func getAdultYears(age int) int {
	return age - 18
}