package main

import "fmt"

func main() {
	//Slice
	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames[1] = "Cathy"
	userNames = append(userNames, "Duncan", "Tonny", "Christine", "Palm")

	// Map
	websites := map[string]string{
		"Google":  "https://google.com",
		"Amazon":  "https://aws.com",
		"Twitter": "https://x.com",
		"IBM":     "https://ibm.com",
	}
	websites["LinkedIn"] = "https://linkedin.com" // Adding an element to a map
	websites["Sony"] = "https://sony.com"
	websites["Huawei"] = "https://huawei.com"
	websites["ubs"] = "https://ubs.com"
	fmt.Println(websites)

	courseRatings := make(map[string]float64, 3)

	courseRatings["Go"] = 4.7
	courseRatings["Java"] = 4.8
	courseRatings["Angular"] = 4.7

	fmt.Println(courseRatings)

	// Using For Loop with Arrays, Slices and Maps
	for i, val := range userNames {
		//...
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", val)
	}

	for key, value := range websites {
		//...
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}

}
