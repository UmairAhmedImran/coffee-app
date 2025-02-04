package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

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
	srv := &http.Server {
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
	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	// TODO: connection to db

	app := &Application {
		Config: cfg,
		// ADD: models later
	}
	
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
