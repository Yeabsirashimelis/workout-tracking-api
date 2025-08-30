package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/app"
	"github.com/Yeabsirashimelis/workout-tracking-api/internal/routes"
)

func main() {
	var port int
	//"port" is the flag name of what is passed from the terminal
 	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close() //means at the very end of execution, then go a head and call this function

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are runnig on %d\n", port)


	err =server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal((err))
	}
}

