package main

import (
	"fmt"
	"go-test/database"
	"go-test/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getDsn() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

// @title			Movie Database
// @version		1.0
// @description	A backend for a Movie Database
// @contact.name	Mikhail Pecherkin
// @contact.email	m.pecherkin.sas@gmail.com
// @BasePath		/
func main() {
	database.SetupDB(getDsn())
	server.Setup(os.Getenv("HOST"))
}
