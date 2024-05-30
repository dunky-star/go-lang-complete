package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"The struggle of a gentleman", "Cute little lion"}
	prices := [10]float64{10.5, 9.2, 11.4, 3.5, 20.1, 6.6, 100.1, 200.5, 30.5, 40.0}
	featuredPrices := prices[2:] // Slice from index to last index
	highlightedPrices := featuredPrices[:1]
	productNames[2] = "A book"
	fmt.Println("List of prices: ", prices)
	fmt.Println("Array of string: ", productNames)
	fmt.Println("Accessing an array element at position no. 3 -> ", prices[2])
	fmt.Println("Sliced from index 2 to last: ", featuredPrices)
	fmt.Println("Sliced from index 0 excluding index 1: ", highlightedPrices)
	fmt.Print("Length of featuredPrices: ", len(featuredPrices), ", Capacity of featuredPrices: ", cap(featuredPrices))
}
