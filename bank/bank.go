package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	// accountBalance := 1000.00
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us at", randomdata.PhoneNumber())
	for {
		switch getMenuChoice() {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			depositAmt := getUserInput("Deposit amount: ")

			if depositAmt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmt
			fmt.Println("Your balance is now", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			withdrawalAmt := getUserInput("Withdrawal amount: ")

			if withdrawalAmt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawalAmt > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmt
			fmt.Println("Your balance is now", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing Go bank!")
			return
		}
	}
}

func getUserInput(message string) (userInput float64) {
	fmt.Printf("%v", message)
	fmt.Scan(&userInput)
	return userInput
}
