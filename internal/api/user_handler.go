package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/Yeabsirashimelis/workout-tracking-api/internal/store"
)



type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger: logger,
	}
}

type registerUserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Bio string `json:"bio"`
}

func (h *UserHandler) ValidateRegisterRequest(req *registerUserRequest)error {
	if req.Username == ""{
		return errors.New("username is required")
	}

	if len(req.Username)> 50{
		return errors.New("username cannot be greater than 50 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

    emailRegex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}$`)
	if !emailRegex.MatchString(req.Email){
		return errors.New("invalid email format")
	}


	if req.Password == ""{
		return errors.New("password is required")
	}

	return nil
}
