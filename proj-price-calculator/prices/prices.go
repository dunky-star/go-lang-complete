package prices

import (
	"bufio"
	"fmt"
	"os"

	"dunky.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Constructor function
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {

	job.loadData() // Loading data read from file

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) loadData() {
	// Reading text data from file
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file.")
		fmt.Println(err)
		return
	}

	var lines []string

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	err = sc.Err()
	if err != nil {
		fmt.Println("Reading the file content failed.")
		file.Close()
		return
	}

	// Converting prices from string read from file to float for calculation
	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Print(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
}
