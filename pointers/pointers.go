package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int 
	agePointer = &age //pointer variable that holds the address of age variable

	fmt.Println("Age", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	*age = *age - 18
	return *age
}