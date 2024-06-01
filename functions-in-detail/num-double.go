package main

import "fmt"

func main() {
	numbers := []int{2, 12, 10, 40, 80, 100}
	doubled := doubbleNumbers(&numbers)
	doubledTransform := transformNumbers(&numbers, double)
	tripleTransform := transformNumbers(&numbers, triple)
	fmt.Println("Doubled numbers func -> ", doubled)
	fmt.Println("Doubled numbers transorm -> ", doubledTransform)
	fmt.Println("Tripled numbers transorm -> ", tripleTransform)

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
