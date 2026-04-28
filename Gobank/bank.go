package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read balance file, defaulting to 1000")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse balance")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		fmt.Println("==============")
		// return
		// panic(err) //we use return or panic to stop the execution of the program when we encounter an error, but panic will also print the stack trace which can be helpful for debugging
		panic("Failed to get balance from file, exiting program")
	}
	fmt.Println("Welcome to Go Bank")

	for {

		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Your account balance: ", accountBalance)

		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var despositAmount float64
			fmt.Scan(&despositAmount)

			if despositAmount <= 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				// return //it stops the execution that no other code is been run
				continue //it stops the execution of current iteration and move to next iteration
			}
			accountBalance += despositAmount
			fmt.Println("Balance updated! New amount ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				return //it stops the execution that no other code is been run
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have")
				return //it stops the execution that no other code is been run
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye! Dude")
			break
		}
	}

	fmt.Println("Thank you for using Go Bank")
}


