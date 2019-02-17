package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/josergdev/go-comments/migration"
	"github.com/josergdev/go-comments/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generate migration to database")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Migration started...")
		migration.Migrate()
		log.Println("Migration done")
	}

	// Init Routes
	router := routes.InitRoutes()

	// Init middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//Init server
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Starting server at https://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Server stopped")
}
