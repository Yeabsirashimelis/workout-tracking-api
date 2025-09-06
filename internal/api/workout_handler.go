package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/store"
	"github.com/Yeabsirashimelis/workout-tracking-api/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger: logger,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
   workoutID, err := utils.ReadIDParam(r)
   if err != nil {
	wh.logger.Printf("ERROR: readIDParam: %v",  err)
   utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout id"})
   }

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)

	if err != nil {
		wh.logger.Printf("ERROR: getWorkoutById: %v",  err)
        utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
         return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout" : workout})

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: decodingCreateWorkout: %v",  err)
        utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
			wh.logger.Printf("ERROR: createWorkout: %v",  err)
           utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}


      utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"workout": createdWorkout})
}


func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request){
	   workoutID, err := utils.ReadIDParam(r)
   if err != nil {
	wh.logger.Printf("ERROR: readIDParam: %v",  err)
   utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout update id"})
   }

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR: getWorkoutById: %v",  err)
        utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if existingWorkout == nil {
	    wh.logger.Printf("ERROR: getWorkoutById: %v",  err)
        utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "can't find workout to update with the specified id"})
		return
	}

	//at this point we are able to find an existing workout
	//we use pointers b/c the zero values of strings and ints are " " and zero, not nil.
	var updateWorkoutRequest struct {
		Title *string `json:"title"`
		Description *string `json:"description"`
		DurationMinutes *int `json:"duration_minutes"`
		CaloriesBurned *int `json:"calories_burned"`
		Entries []store.WorkoutEntry `json:"entries"`
}
 err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)
 if err != nil {
	 wh.logger.Printf("ERROR: decodingUpdateRequest: %v",  err)
     utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request payload"})
	return
 }

 if updateWorkoutRequest.Title != nil {
	existingWorkout.Title = *updateWorkoutRequest.Title
 }
  if updateWorkoutRequest.Description != nil {
	existingWorkout.Description = *updateWorkoutRequest.Description
 }
  if updateWorkoutRequest.DurationMinutes != nil {
	existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
 }
  if updateWorkoutRequest.CaloriesBurned != nil {
	existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
 }
  if updateWorkoutRequest.Entries != nil {
	existingWorkout.Entries = updateWorkoutRequest.Entries
 }

 err = wh.workoutStore.UpdateWorkout(existingWorkout)
 if err != nil {
	 wh.logger.Printf("ERROR: updatingWorkout: %v",  err)
     utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
	return 
 }

 utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": existingWorkout})

}

func (wh *WorkoutHandler) HandleDeleteWorkoutByID(w http.ResponseWriter, r *http.Request) {
	
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
	    wh.logger.Printf("ERROR: readIDParam: %v",  err)
        utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout delete id"})
	}

	err = wh.workoutStore.DeleteWorkout(workoutID)
	if err == sql.ErrNoRows {
		  wh.logger.Printf("ERROR: deletingWorkout: %v", err)
          utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "workout not found to delete with sepcified id"})
	    return
	}

	if err != nil {
		  wh.logger.Printf("ERROR: deletingWorkout: %v", err)
          utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message":"workout deleted successfully"})
}