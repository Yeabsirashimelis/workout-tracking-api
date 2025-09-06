package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}


	//run the migrations for our test db
	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("migrating test db error: %v", err)
	}

	//clear the database after any testing to it
	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("truncating tables %v", err)
	}

	return db
}

/*
    ANONYMOUS STRUCTS
[] before a struct means “this is a slice of these structs”.
Anonymous structs are often used with slices for quick temporary collections without creating a named struct type.
*/

func TestCreateworkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name string
		workout *Workout
		wantErr bool
	} {
		{
			name: "valid workout",
			workout: &Workout{
				Title: "push day",
				Description: "upper body day",
				DurationMinutes: 60,
				CaloriesBurned: 200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "bench press",
						Sets: 3,
						Reps: Intptr(10),
						Weight: FloatPointer(135.5),
						Notes: "warm up properly",
						OrderIndex: 1,

					},
				},
		}, wantErr: false,
	
      }, {
		name: "workout with invalid entries",
		workout: &Workout{
			Title: "full body",
				Description: "complete workout",
				DurationMinutes: 60,
				CaloriesBurned: 200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "plank",
						Sets: 3,
						Reps: Intptr(10),
						Notes: "warm up properly",
						OrderIndex: 1,

					},
					{
						ExerciseName: "squats",
						Sets: 4,
						Reps: Intptr(12),
						Weight: FloatPointer(185.5),
						DurationSeconds: Intptr(60),
						Notes: "full depth",
						OrderIndex: 2,

					},
				},
		},
		wantErr: true,
	  },  
   } 

   for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
		createdWorkout, err := store.CreateWorkout(tt.workout)
		if tt.wantErr {
			assert.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, tt.workout.Title, createdWorkout.Title)
		assert.Equal(t, tt.workout.Description, createdWorkout.Description)
		assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)

		retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
		require.NoError(t, err)

		assert.Equal(t, createdWorkout.ID, retrieved.ID)
		assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

		for i, _ := range retrieved.Entries {
			assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
			assert.Equal(t, tt.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
			assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)
			assert.Equal(t, tt.workout.Entries[i].Notes, retrieved.Entries[i].Notes)
		}
	  })
   }
}


func Intptr(i int) *int {
	return &i
}

func FloatPointer(i float64) *float64 {
	return &i
}