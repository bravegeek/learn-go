package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	var revenue, expenses, taxRate float64
	// var earnings_before_tax, earnings_after_tax, ratio float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	// fmt.Print("EBT: ")
	fmt.Println("EBT: ", ebt)
	// fmt.Print("Profit: ")
	fmt.Println("Profit: ", profit)
	// fmt.Print("Ratio: ")
	fmt.Println("Ratio: ", ratio)
}
