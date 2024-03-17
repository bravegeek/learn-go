package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
