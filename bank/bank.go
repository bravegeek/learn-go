package main

import "fmt"

func main() {
	accountBalance := 1000.00

	fmt.Println("Welcome to Go Bank!")
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
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing Go bank!")
			return
		}
	}
}

func getMenuChoice() (menuChoice int) {
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

func processUserInput(menuChoice int) {

}
