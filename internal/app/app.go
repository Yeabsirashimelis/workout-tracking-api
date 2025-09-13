package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/api"
	"github.com/Yeabsirashimelis/workout-tracking-api/internal/middleware"
	"github.com/Yeabsirashimelis/workout-tracking-api/internal/store"
	"github.com/Yeabsirashimelis/workout-tracking-api/migrations"
)




type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	UserHandler *api.UserHandler
	TokenHandler *api.TokenHandler
    Middleware middleware.UserMiddleware
	DB *sql.DB
}

//method on the application struct
func NewApplication() (*Application, error){
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout,"",log.Ldate | log.Ltime )


//our stores will go here
workoutStore := store.NewPostgresWorkoutStore(pgDB)
userStore := store.NewPostgresUserStore(pgDB)
tokenStore := store.NewPostgresTokenStore(pgDB)


//our handlers will go here

workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
userHandler := api.NewUserHandler(userStore, logger)
tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
middlewareHandler := middleware.UserMiddleware{Userstore: userStore}

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		UserHandler: userHandler,
		TokenHandler: tokenHandler,
		Middleware: middlewareHandler,
		DB: pgDB,
	}
return app, nil
}

//passing things to params by value will copy it entirely, the request here is a pointer b/c of like that
// b/c copying large requests will be inefficient
func(a *Application) HealthCheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Status is available\n")
}
