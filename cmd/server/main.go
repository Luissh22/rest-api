package main

import (
	transportHTTP "github.com/luissh22/rest-api/internal/transport/http"
	"log"
	"net/http"
)

type App struct {

}

func (a *App) Run() error {
	log.Println("Running application")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
