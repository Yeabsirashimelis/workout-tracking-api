package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/store"
	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64) //base 10 and 64 bit int

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to created workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		http.Error(w, "failed to created workout", http.StatusInternalServerError)
		return
	}


	fmt.Println("....................workout successfully created......................")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}
