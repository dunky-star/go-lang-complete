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

	fmt.Printf("The EBITDA: %.2f\n", ebitda)
	fmt.Printf("The Profit: %.2f\n", profit)
	fmt.Printf("The Ratio: %.2f\n", ratio)

}