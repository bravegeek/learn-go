package main

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

func main() {

	productNames := [4]string{"A book"}

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[0])

	featuredPrices := prices[1:3] //1st number is inclusive, 2nd is exclusive
	fmt.Println(featuredPrices)
}
