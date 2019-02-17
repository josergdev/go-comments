package routes

import (
	"github.com/gorilla/mux"
	"github.com/josergdev/go-comments/controllers"
)

// SetLoginRouter set route for Login handler
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
