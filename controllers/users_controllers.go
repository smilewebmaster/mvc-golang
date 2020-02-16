package controllers

import (
	"encoding/json"
	"github.com/dmolina79/mvc-golang/utils"
	"net/http"
	"strconv"

	"github.com/dmolina79/mvc-golang/services"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonUserErr, _:= json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonUserErr)
		return
	}

	user, apiError := services.GetUser(userID)

	if apiError != nil {
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(apiError.Message))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
