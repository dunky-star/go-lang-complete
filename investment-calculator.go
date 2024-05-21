package main // This is the main entry point (package main)

import (
	"fmt"
	"math"
)

func main(){

	const inflationRate = 2.5 // Can not be reassigned
	// Assigning type explicitly to override inferred value
	var investmentAmount float64 = 0.0
	expectedReturnRate := 5.5 
	var years float64

	fmt.Print("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years to invest: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := investmentAmount * math.Pow(1 + inflationRate/100, years)
    fmt.Print(`The investment return is: `, futureValue, "\n")
	fmt.Print(`The real value after inflation: `, futureRealValue)
}