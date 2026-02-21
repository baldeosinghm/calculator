package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates)) // initialize list to hold channels
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool) // Create new channels, one per index
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)   // Either filemanager or cmdmanager's methods are acceptable here
		go priceJob.Process(doneChans[index], errorChans[index]) // Give each function call to Process() a channel

		// if err != nil {
		// 	fmt.Println("Could not process job.")
		// 	fmt.Println(err)
		// }
	}

	// The "select" stament: A control structure meant to be used with channels. Allows us
	// to listen for data from multiple channels. Like the switch stmt, "select" won't
	// wait for the next channel to finish if one comes back with an output.

	// I'll do a for loop on taxRates and run my channels in there; this approach on the
	// taxRates themselves is possible. I could also place the select statement into a
	// for loop with all the channels I do have; so either doneChans or errorChans. But,
	// I can just stick with taxRates b/c it's used as a source by both channels.

	for index := range taxRates {
		select {
		case err := <-errorChans[index]: // First case is if we catch an error in our errors channel
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]: // Second case is for each price calculated
			fmt.Println("Done!")
		}
	}

	// // Tell Go to wait until every function call finishes before exiting program
	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
