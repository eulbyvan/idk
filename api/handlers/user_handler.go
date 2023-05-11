/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 11 2023 10:19:37 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

// UserHandler defines the HTTP handlers for the user entity
type UserHandler struct {
	userService domain.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUserByIDHandler retrieves a user by their ID
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		// Handle the error and write the response
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	// Marshal the user to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		// Handle the error and write the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// CreateUserHandler creates a new user
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a User struct
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle the error and write the response
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Create the user
	err = h.userService.CreateUser(&user)
	if err != nil {
		// Handle the error and write the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Write the response
	w.WriteHeader(http.StatusCreated)
}
