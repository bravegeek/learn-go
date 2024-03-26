package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeFinancialsToFile(ebt, profit, ratio)

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

	userInput, err := validateUserInput(userInput)

	if err != nil {
		fmt.Println(err)
		panic("aaaaah!")
	}

	return userInput
}

func validateUserInput(userInput float64) (float64, error) {
	if userInput < 0 || userInput == 0 {
		return userInput, errors.New("value must be greater than 0")
	}
	return userInput, nil
}

const financialsFile = "financials.txt"

func writeFinancialsToFile(ebt float64, profit float64, ratio float64) {

	stringToWrite := []byte(fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio))
	os.WriteFile(financialsFile, stringToWrite, 0644)
}
