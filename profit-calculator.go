// package main

// import "fmt"

// func main(){

// 	//var income float64
// 	//var expenses float64
// 	// var taxRate float64

// 	income := getUserInput("Income: ")
// 	expenses := getUserInput("Expenses: ")
// 	taxRate := getUserInput("Tax Rate: ")

// 	ebitda, profit, ratio := calculateFinancials(income, expenses, taxRate)

// 	fmt.Printf("The EBITDA: %.2f\n", ebitda)
// 	fmt.Printf("The Profit: %.2f\n", profit)
// 	fmt.Printf("The Ratio: %.3f\n", ratio)

// }

// func calculateFinancials(income, expenses, taxRate float64) (float64, float64, float64) {
// 	ebitda := income - expenses;
// 	profit := ebitda - (1 - taxRate/100)
// 	ratio := ebitda / profit
//     return ebitda, profit, ratio
// }

// func getUserInput (infoText string) float64{
// 	var userInput float64
// 	fmt.Print(infoText)
// 	fmt.Scan(&userInput)
// 	return userInput
// }
