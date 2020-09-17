package handlers

import (
	"github.com/bpsaunders/user-api/service"
	"github.com/gorilla/mux"
	"net/http"
)

// Register registers handler functions against all available routes
func Register(router *mux.Router, userService service.UserService) {

	router.HandleFunc("/health-check", healthCheck)
	router.Handle("/users", NewCreateUserHandler(userService)).Methods(http.MethodPost)
	router.Handle("/users", NewGetAllUsersHandler(userService)).Methods(http.MethodGet)
	router.Handle("/users/{user_id}", NewGetUserHandler(userService)).Methods(http.MethodGet)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
