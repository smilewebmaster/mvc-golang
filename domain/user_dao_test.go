package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a valid user with id 0")
	assert.NotNil(t, err, "Expected an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 status code")
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserGoGood(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Doug", user.FirstName)
	assert.EqualValues(t, "Molina", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}