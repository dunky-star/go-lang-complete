package main

import "fmt"

func main() {
	numbers := []int{2, 12, 10, 40, 80, 100}
	doubled := doubbleNumbers(&numbers)
	fact := numFactorial(5)
	rfact := recursiveFactorial(5)
	quadTransform := transformNumbers(&numbers, func(number int) int { return number * 4 }) // Anonymous function
	doubleClosure := createTransformer(2)
	tripleTransformer := createTransformer(3)
	doubledTransform := transformNumbers(&numbers, doubleClosure)
	tripleTransform := transformNumbers(&numbers, tripleTransformer)
	fmt.Println("Doubled numbers func -> ", doubled)
	fmt.Println("Doubled numbers transorm -> ", doubledTransform)
	fmt.Println("Tripled numbers transorm -> ", tripleTransform)
	fmt.Println("Quad numbers from Anonymous -> ", quadTransform)
	fmt.Println("Factorial without recursion -> ", fact)
	fmt.Println("Factorial with recursion -> ", rfact)
	fmt.Println("The sum of array elements  in numbers = ", sumUp(numbers))
	fmt.Println("The sum of any elements = ", sumArbitrary(10, 1000, 39, 40, 20, 1))
	anotherSum := sumArbitrary(numbers...) // Splitting slcies into parameter values.
	fmt.Println("Splitting slices and summing up= ", anotherSum)
}

// Write a function to double the int values in a slice and return the doubled slice.
func doubbleNumbers(numbers *[]int) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}
	return dNumbers
}

// Passing functions as parameters value.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

// Utility func 1
func double(number int) int {
	return number * 2
}

// Utility func 2
func triple(number int) int {
	return number * 3
}

// Closure concept -> Factory method
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Claculating factorial without recursion
func numFactorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

// Recursion  -> Calculating factorial of a number with recursion
func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

// Write a function to sum up all the integer array elements
func sumUp(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

// Variadic function to sum up arbitrary elements
func sumArbitrary(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}
