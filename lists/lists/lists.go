package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} // prices is really a slice, go handles the backing array
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 7.11)
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 5.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {

// 	productNames := [4]string{"A book"}

// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[0])

// 	featuredPrices := prices[1:3] //1st number is inclusive, 2nd is exclusive

// 	highlightedPrices := featuredPrices[1:] // slice of a slice

// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)

// 	//slices are a window (pointer) into an array
// 	featuredPrices[0] = 199.99 //modify a value from the slice
// 	fmt.Println(prices)        // orig array

// 	// startingSlice := prices[:3]
// 	// fmt.Println(startingSlice)

// 	// endingSlice := prices[1:]
// 	// fmt.Println(endingSlice)

// 	fmt.Println(len(featuredPrices), cap(featuredPrices)) // length and capacity
// }
