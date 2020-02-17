package services

import (
	"github.com/dmolina79/mvc-golang/domain"
	"github.com/dmolina79/mvc-golang/utils"
)

type usersService struct {

}

var (
	UsersService usersService
)

// GetUser ...
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
