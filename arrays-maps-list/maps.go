package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	websites["LinkedIn"] = "https://linkedin.com" // Adding an element to a map
	fmt.Println(websites)
}
