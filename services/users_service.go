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
	return domain.UserDao.GetUser(userID)
}
