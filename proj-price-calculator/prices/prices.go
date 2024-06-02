package prices

import (
	"fmt"

	"dunky.com/price-calculator/conversion"
	"dunky.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// Constructor function
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
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

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) loadData() {
	// Reading text data from file
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Print(err)
		return
	}

	// Converting prices from string read from file to float for calculation
	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Print(err)
		return
	}

	job.InputPrices = prices
}
