package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank")

	for  {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
				return //it stops the execution that no other code is been run
			}
			accountBalance += despositAmount
			fmt.Println("Balance updated! New amount ", accountBalance)
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
		} else {
			fmt.Println("Goodbye! Dude")
		}
	}

}
