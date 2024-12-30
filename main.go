package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	//create a slice of prices
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	//we then create a for loop that loops through the tax rates and calculate the tax rate of each prices

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncudedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
		// //we create another nested for loop that goes through the prices for every tax rate, then we alculate new prices based on the taxRate we got
		// //in the current loop iteration
		// taxIncludedPrices := make([]float64, len(prices)) //we use the make function to create a new slice, unlike the map we would specify the length it should contain
		// for priceIndex, price := range prices {
		// 	taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		// }
		// result[taxRate] = taxIncludedPrices
	}

}
