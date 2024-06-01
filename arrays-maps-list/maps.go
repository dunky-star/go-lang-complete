package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	websites["LinkedIn"] = "https://linkedin.com" // Adding an element to a map
	fmt.Println(websites)

	courseRatings := make(map[string]float64, 3)

	courseRatings["Go"] = 4.7
	courseRatings["Java"] = 4.8
	courseRatings["Angular"] = 4.7

	fmt.Print(courseRatings)

}
