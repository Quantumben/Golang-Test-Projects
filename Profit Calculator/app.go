package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
//1. Validate user input
// show error message and exit if invalid input is given
//no negative numbers
//not even 0
//Store calculated results into file

func main() {
	revenue, err := getUserInput("Enter the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter the expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	tax_rate, err := getUserInput("Enter the tax rate (as a decimal): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("result.txt",[]byte(results),0644)
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := profit / revenue
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	// validate user input
	if userInput <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return userInput, nil
}
