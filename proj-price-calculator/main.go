package main

import (
	"fmt"

	//"dunky.com/price-calculator/cmdmanager"
	"dunky.com/price-calculator/filemanager"
	"dunky.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates)) // goroutines and channels

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])
		// if err != nil {
		// 	fmt.Println("couldn't process job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}

}
