package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dmolina79/mvc-golang/services"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
