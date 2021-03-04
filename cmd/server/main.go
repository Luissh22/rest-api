package main

import "log"

type App struct {

}

func (a *App) Run() error {
	log.Println("Running application")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
