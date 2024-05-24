package main

import "fmt"

func main() {

	fmt.Println("\nWELCOME TO DUNKY BANK")
	fmt.Println("=====================\n")

	// For Loop is the only kind of Loop in Go but it's also flexible.
	for { // -> Infinite Loop in Go

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit\n")

		var accountBalance float64 = 4000

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Balance: ", accountBalance, "\n")
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Print("Balance updated! New bal: ", accountBalance, "\n")

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance -= withdrawAmount
			if accountBalance <= 0 {
				fmt.Print("You have insufficient account balance\n")
				continue
			} else {
				fmt.Print("Withdraw successful! New bal: ", accountBalance, "\n")

			}

		default:
			fmt.Print("Goodbye! ")
			fmt.Print("Thanks for choosing our bank")
			return
			// break // -> To allow code outside of the loop to execute when using if condition

		} // End of Switch statement

	} // End of For Loop

}
