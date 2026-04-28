package main

import "fmt"

func main(){

	var revenue,expenses,tax_rate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate (as a decimal): ")
	fmt.Scan(&tax_rate)

	profit := revenue - expenses * tax_rate
	fmt.Print(profit)
}