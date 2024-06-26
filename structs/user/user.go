package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	created   time.Time
}

type Admin struct {
	email    string
	password string
	User     // anonymous embedding
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// mutation function needs a pointer to the struct
// otherwise it receives a copy
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("missing values")
	}

	// returns a pointer to the user created
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		created:   time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			created:   time.Now(),
		},
	}
}
