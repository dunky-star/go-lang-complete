package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalFile = "balance.txt"

func main() {

	fmt.Println("\nWELCOME TO DUNKY BANK")
	fmt.Println("=====================\n")

	// For Loop is the only kind of Loop in Go but it's also flexible.
	for { // -> Infinite Loop in Go

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit\n")

		var accountBalance, err = getBalanceFromFile()

		if err != nil {
			fmt.Println("Error")
			fmt.Println(err)
			fmt.Print("----------")
			panic("Sorry, can't continue") // To terminate upon encountering an error
		}

		var choice int
		fmt.Print("\nYour choice: ")
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
			writeBalanceToFile(accountBalance)

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
				writeBalanceToFile(accountBalance)
			}

		default:
			fmt.Print("Goodbye! ")
			fmt.Print("Thanks for choosing our bank")
			return
			// break // -> To allow code outside of the loop to execute when using if condition

		} // End of Switch statement

	} // End of For Loop

} // End of main()

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalFile)

	if err != nil {
		return 0, errors.New("Failed to read file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}
