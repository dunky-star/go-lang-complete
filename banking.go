package main

import "fmt"

func main() {

	// For Loop is the only kind of Loop in Go but it's also flexible.
	for { // -> Infinite Loop in Go
		fmt.Println("\nWELCOME TO DUNKY BANK")
		fmt.Println("=====================\n")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit\n")

		var accountBalance float64 = 0

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Balance: ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0")
				return
			}
			accountBalance += depositAmount
			fmt.Print("Balance updated! New bal: ", accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0")
			}
			accountBalance -= withdrawAmount
			if accountBalance <= 0 {
				fmt.Print("You have insufficient account balance")
				return
			} else {
				fmt.Print("Withdraw successful! New bal: ", accountBalance)

			}

		} else {
			fmt.Print("Goodbye!")
		}
	}
}
