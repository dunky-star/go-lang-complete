package user

import (
	"errors"
	"fmt"
	"time"
)

// struct: custom data type
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Struct embedding => equivalent to inheritance in Java, etc
type Admin struct {
	email    string
	password string
	User     User // Embedding
}

// Constructor function
func New(firstName, lastName, birthDate string) (*User, error) {
	// Constructor helps us to perform validation
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil

}

// Constructor function
func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Duncan",
			lastName:  "Kaligs",
			birthDate: "____",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetail() { // To pass a pointer to the function
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}
