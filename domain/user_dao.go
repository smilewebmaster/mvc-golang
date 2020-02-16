package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Doug", LastName: "Molina", Email: "myemail@gmail.com"},
	}
)

// GetUser ...
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))
}
