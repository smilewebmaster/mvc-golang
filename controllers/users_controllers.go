package controllers

import (
	"github.com/dmolina79/mvc-golang/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/dmolina79/mvc-golang/services"
)



// GetUser ...
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)

		return
	}

	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK,user)
}
