package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/api"
	"github.com/Yeabsirashimelis/workout-tracking-api/internal/store"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}

//method on the application struct
func NewApplication() (*Application, error){
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout,"",log.Ldate | log.Ltime )

//our store will go here

//our handlers will go here
workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}
return app, nil
}

//passing things to params by value wil/ copy it entirely, the request here is a pointer b/c of like that
// b/c copying large requests will be inefficient
func(a *Application) HealthCheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Status is available\n")
}