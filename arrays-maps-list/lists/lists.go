package lists

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Motorsport", "Dancing", "Scientific Journals"}
	prices := []float64{10.5, 9.2, 11.4, 3.5, 20.1, 6.6, 100.1, 200.5, 30.5, 40.0}
	discountPrices := []float64{1.5, 2.2, 4.5, 6.6}    // Dynamic arrays
	updatedPrices := append(prices, discountPrices...) // Unpacking list
	featuredPrices := prices[2:]                       // Slice from index to last index
	highlightedPrices := featuredPrices[:1]
	hobbies[2] = "Read books"
	fmt.Println("List of prices: ", prices)
	fmt.Println("Hobbies: ", hobbies)
	fmt.Println("Accessing an array element at position no. 3 -> ", prices[2])
	fmt.Println("Sliced from index 2 to last: ", featuredPrices)
	fmt.Println("Sliced from index 0 excluding index 1: ", highlightedPrices)
	fmt.Println("Length of featuredPrices: ", len(featuredPrices), ", Capacity of featuredPrices: ", cap(featuredPrices))
	fmt.Println("Dynamic arrays: ", updatedPrices)

	// Array of a struct
	product := []Product{
		{
			"first_product",
			"Coffee beans",
			99.9,
		},
		{
			"second_product",
			"Beautiful cup",
			100.9,
		},
		{
			"third_product",
			"Shovel Robot",
			1000.8,
		},
	}
	fmt.Println("Product List: ", product)

	newProduct := Product{
		"fourth_product",
		"Sweeping broom",
		40.5,
	}

	product = append(product, newProduct)
	fmt.Println("Product after append: ", product)

} // End of main() func
