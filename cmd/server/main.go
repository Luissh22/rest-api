package main

import (
	"github.com/gorilla/mux"
	"github.com/luissh22/rest-api/internal/comment"
	httpConstants "github.com/luissh22/rest-api/internal/constants/http"
	"github.com/luissh22/rest-api/internal/database/postgres"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) Run() error {
	log.Println("Running application")

	// Setup database
	db, err := postgres.NewDatabase()

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&comment.Comment{})

	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.Use(jsonMiddleware)

	// Setup comments handler
	commentsService := comment.NewService(db)
	commentHandler := comment.NewHandler(r, commentsService)
	commentHandler.SetupRoutes()

	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Println("Failed to setup server")
		return err
	}

	return nil
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(httpConstants.ContentType, httpConstants.ApplicationJSON)
		next.ServeHTTP(w, r)
	})
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
