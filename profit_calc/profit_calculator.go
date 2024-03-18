package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	var revenue, expenses, taxRate float64
	// var earnings_before_tax, earnings_after_tax, ratio float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue = getUserInput("Revenue")

	expenses = getUserInput("Expenses")

	taxRate = getUserInput("Tax Rate")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	// fmt.Print("EBT: ")
	fmt.Println("EBT: ", ebt)
	// fmt.Print("Profit: ")
	fmt.Println("Profit: ", profit)
	// fmt.Print("Ratio: ")
	fmt.Println("Ratio: ", ratio)
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
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
