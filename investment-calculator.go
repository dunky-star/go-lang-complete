// package main // This is the main entry point (package main)

// import (
// 	"fmt"
// 	"math"
// )

// const inflationRate = 2.5 // Can not be reassigned

// func main(){

	
// 	// Assigning type explicitly to override inferred value
// 	var investmentAmount float64
// 	expectedReturnRate := 5.5 
// 	var years float64

// 	fmt.Print("Investment amount: ")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Print("Expected return rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	fmt.Print("Years to invest: ")
// 	fmt.Scan(&years)

// 	futureValue, futureRealValue  := calculateFutureValues(investmentAmount, expectedReturnRate, years)
// 	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
// 	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
//     fmt.Printf("The investment return is: %.2f \n", futureValue)
// 	fmt.Printf("The real value after inflation: %.2f", futureRealValue)
// }

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64){
// 	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1 + inflationRate/100, years)
// 	return fv, rfv
// } 
