package main

import (
	"flag"
	"log"

	"github.com/josergdev/go-comments/migration"
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
}
