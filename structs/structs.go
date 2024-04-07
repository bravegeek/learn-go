package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	created   time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	// func newUser(firstName, lastName, birthDate string) user {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Missing values")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		created:   time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date: ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthDate,
	// 	created:   time.Now(),
	// }

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }
