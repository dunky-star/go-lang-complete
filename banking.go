package main

import (
	"fmt"
	"time"

	"dunky.com/banking/file_utility"
	"github.com/Pallinder/go-randomdata"
)

// struct: custom data type
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

const accountBalFile = "balance.txt"

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	fmt.Println("\nWELCOME TO DUNKY BANK")
	fmt.Println("Reach us 24/7 on ", randomdata.PhoneNumber())
	fmt.Println("================================\n")
	outputUserDetail(&appUser) // Pointer to appUser

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

func outputUserDetail(u *user) { // To pass a pointer to the function
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
