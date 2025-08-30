package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r,"id")
	if paramsWorkoutID == ""{
		http.NotFound(w, r)
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)//base 10 and 64 bit int
	
	if err != nil {
		http.NotFound(w, r)
		return
	}
 
	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "created a workout\n")
	
}