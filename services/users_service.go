package services

import "github.com/dmolina79/mvc-golang/domain"

// GetUser ...
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
