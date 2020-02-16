package services

import (
	"github.com/dmolina79/mvc-golang/domain"
	"github.com/dmolina79/mvc-golang/utils"
)

// GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
