package app

import (
	"net/http"

	"github.com/dmolina79/mvc-golang/controllers"
)

// StartApp Start...
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
