package services

import (
	"github.com/dmolina79/mvc-golang/domain"
	"github.com/dmolina79/mvc-golang/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type usersDaoMock struct {

}

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUsersService_UserNotFoundInDB(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exist",
			StatusCode: http.StatusNotFound,
			Code:       "404",
		}
	}
	user, err:= UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestUsersService_GetUser(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
				ID:        123,
				FirstName: "Homer",
				LastName:  "Simpson",
				Email:     "hsimpson@gmail.com",
			}, nil
	}

	u, err:= UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.EqualValues(t, 123, u.ID)
	assert.EqualValues(t, "Homer", u.FirstName)
	assert.EqualValues(t, "Simpson", u.LastName)
	assert.EqualValues(t, "hsimpson@gmail.com", u.Email)
}
