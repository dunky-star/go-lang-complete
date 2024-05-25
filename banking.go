package main

import (
	"fmt"

	"dunky.com/banking/file_utility"
	"github.com/Pallinder/go-randomdata"
)

const accountBalFile = "balance.txt"

func main() {

	fmt.Println("\nWELCOME TO DUNKY BANK")
	fmt.Println("Reach us 24/7 on ", randomdata.PhoneNumber())
	fmt.Println("================================\n")

	// For Loop is the only kind of Loop in Go but it's also flexible.
	for { // -> Infinite Loop in Go

		presentOptions()

		var accountBalance, err = file_utility.GetFloatFromFile(accountBalFile)

		if err != nil {
			fmt.Println("Error")
			fmt.Println(err)
			fmt.Print("----------")
			//panic("Can't continue, sorry.") // To terminate upon encountering an error
		}

		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("\nBalance: ", accountBalance, "\n")
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
			file_utility.WriteFloatToFile(accountBalance, accountBalFile)

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance -= withdrawAmount
			if accountBalance <= 0 {
				fmt.Println("You have insufficient account balance\n")
				continue
			} else {
				fmt.Print("Withdraw successful! New bal: ", accountBalance, "\n")
				file_utility.WriteFloatToFile(accountBalance, accountBalFile)
			}

		default:
			fmt.Println("Goodbye! ")
			fmt.Println("Thanks for choosing our bank")
			return
			// break // -> To allow code outside of the loop to execute when using if condition

		} // End of Switch statement

	} // End of For Loop

} // End of main()

func presentOptions() {
	fmt.Println("\n1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit\n")
}
