package main // This is the main entry point (package main)

import (
	"fmt"
	"math"
)

func main(){

	var investmentAmount float64 = 1000 // Assigning type explicitly to override inferred value
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
    fmt.Print(`The investment return is: `, futureValue)
}