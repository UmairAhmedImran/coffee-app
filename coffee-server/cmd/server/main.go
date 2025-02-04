package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UmairAhmedImran/coffee-app/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	port := os.Getenv("PORT")
	fmt.Println("API is listening on port", port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO: add router
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DNS")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	// TODO: connection to db

	app := &Application{
		Config: cfg,
		// ADD: models later
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
