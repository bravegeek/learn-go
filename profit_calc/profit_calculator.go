package main

import (
	"fmt"
)

func main() {
	// var revenue, expenses, taxRate float64
	// var earnings_before_tax, earnings_after_tax, ratio float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(message string) (userInput float64) {
	fmt.Printf("%v: ", message)
	fmt.Scan(&userInput)
	return userInput
}
