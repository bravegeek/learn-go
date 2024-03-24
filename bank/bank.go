package main

import "fmt"

func main() {
	accountBalance := 1000.00
	menuChoice := getMenuChoice()

	if menuChoice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if menuChoice == 2 {
		depositAmt := getUserInput("Deposit amount: ")

		if depositAmt <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmt
		fmt.Println("Your balance is now", accountBalance)
	} else if menuChoice == 3 {
		withdrawalAmt := getUserInput("Withdrawal amount: ")

		if withdrawalAmt <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}
		if withdrawalAmt > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withdrawalAmt
		fmt.Println("Your balance is now", accountBalance)
	} else {
		println("Goodbye!")
	}
}

func getMenuChoice() (menuChoice int) {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("\nYour choice: ")
	fmt.Scan(&menuChoice)
	return menuChoice
}

func getUserInput(message string) (userInput float64) {
	fmt.Printf("%v", message)
	fmt.Scan(&userInput)
	return userInput
}
