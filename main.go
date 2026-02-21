package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates)) // initialize list to hold channels

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool) // Create new channels, one per index
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate) // Either filemanager or cmdmanager's methods are acceptable here
		go priceJob.Process(doneChans[index])                  // Give each function call to Process() a channel

		// if err != nil {
		// 	fmt.Println("Could not process job.")
		// 	fmt.Println(err)
		// }
	}

	// Tell Go to wait until every function call finishes before exiting program
	for _, doneChan := range doneChans {
		<-doneChan
	}
}
