package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"The struggle of a gentleman", "Cute little lion"}
	prices := [6]float64{10.5, 9.2, 11.4, 3.5, 20.1, 6.6}
	fmt.Println("List of prices: ", prices)
	fmt.Println("Array of string: ", productNames)
	fmt.Println("Accessing an array element at position no. 3 -> ", prices[2])
}
