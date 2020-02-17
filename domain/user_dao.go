package domain

import (
	"fmt"
	"github.com/dmolina79/mvc-golang/utils"
	"net/http"
)

type userDao struct {}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}


var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Doug", LastName: "Molina", Email: "myemail@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// GetUser ...
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
