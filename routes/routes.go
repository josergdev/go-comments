package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes init all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)

	return router
}
