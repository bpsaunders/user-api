package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bpsaunders/user-api/models"
	"github.com/bpsaunders/user-api/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// CreateUserHandler offers a handler by which to create a user
type CreateUserHandler struct {
	service service.UserService
}

// NewCreateUserHandler returns a new CreateUserHandler
func NewCreateUserHandler(service service.UserService) CreateUserHandler {
	return CreateUserHandler{
		service,
	}
}

// GetUserHandler offers a handler by which to fetch a user
type GetUserHandler struct {
	service service.UserService
}

// NewGetUserHandler returns a new GetUserHandler
func NewGetUserHandler(service service.UserService) GetUserHandler {
	return GetUserHandler{
		service,
	}
}

// GetAllUsersHandler offers a handler by which to fetch all users
type GetAllUsersHandler struct {
	service service.UserService
}

// NewGetAllUsersHandler returns a new GetAllUsersHandler
func NewGetAllUsersHandler(service service.UserService) GetAllUsersHandler {
	return GetAllUsersHandler{
		service,
	}
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to decode request body to user struct: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debug(
		fmt.Sprintf(
			"Submitted user - first name: %s, last name: %s, email: %s, country: %s",
			user.FirstName, user.LastName, user.Email, user.Country))

	responseType, validationErrors, err := h.service.CreateUser(&user)

	if responseType == service.Error {
		log.Error(fmt.Sprintf("Error encountered when creating user: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if responseType == service.Conflict {
		log.Info("Attempt made to create a user that already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	if responseType == service.InvalidData {
		log.Info("Invalid data submission")
		log.Debug(fmt.Sprintf("errors returned: %s", validationErrors))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(validationErrors)
		if err != nil {
			log.Error(fmt.Sprintf("Error writing response: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Info("User created successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID := vars["user_id"]
	if userID == "" {
		log.Info("No userID in url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseType, user, err := h.service.GetUser(userID)
	if responseType == service.Error {
		log.Error(fmt.Sprintf("Error encountered when fetching user: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if responseType == service.NotFound {
		log.Info("User not found")
		log.Debug(fmt.Sprintf("User not found by id: %s", userID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Info("User fetched successfully")
	log.Debug(fmt.Sprintf("User found with id: %s", userID))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h GetAllUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	responseType, users, err := h.service.GetAllUsers()
	if responseType == service.Error {
		log.Error(fmt.Sprintf("Error encountered when fetching users: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info("Users fetched successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
