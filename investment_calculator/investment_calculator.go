package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years, expectedReturnRate float64
	// years, expectedReturnRate := 10.0, 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
	// fmt.Printf("Future Value: %.0f\nFuture Real Value: %.1f\n", futureValue, futureRealValue)
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
}
