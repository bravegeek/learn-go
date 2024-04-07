package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date: ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("me@me.com", "secure123")
	admin.AdminUser.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
