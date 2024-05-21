package main

import "fmt"

func main(){

	var income float64
	var expenses float64
	var taxRate float64

	fmt.Print("Income: ")
	fmt.Scan(&income)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

    ebitda := income - expenses;
	profit := ebitda - (1 - taxRate/100)
	ratio := ebitda / profit

	fmt.Print(`The EBITDA: `, ebitda, "\n")
	fmt.Print(`The Profit: `, profit, "\n")
	fmt.Print(`The Ratio: `, ratio, "\n")

}