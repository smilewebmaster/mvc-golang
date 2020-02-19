package app

import (
	"github.com/dmolina79/mvc-golang/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}