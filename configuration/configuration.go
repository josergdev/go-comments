package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Driver of mysql to return a conenction to mysql database
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Configuration model for config.json
type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration read config.json and parse it to Configuration struct
func getConfiguration() configuration {
	var config configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

// GetConnection returns connection to the configured database
func GetConnection() *gorm.DB {
	c := getConfiguration()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
