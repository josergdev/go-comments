package routes

import (
	"github.com/gorilla/mux"
	"github.com/josergdev/go-comments/controllers"
	"github.com/urfave/negroni"
)

// SetUserRouter set route for Register handler
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.Register).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(negroni.Wrap(subRouter)),
	)
}
