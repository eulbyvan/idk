/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 12:41:11 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"log"
	"net/http"

	"github.com/eulbyvan/idk/app/go-user-management/api/handlers"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
	"github.com/eulbyvan/idk/app/go-user-management/internal/repository"
)

func main() {
	// Create the user repository
	userRepository := repository.NewUserRepositoryImpl()

	// Create the user service
	userService := domain.NewUserService(userRepository)

	// Create the user handler
	userHandler := handlers.NewUserHandler(*userService)

	// Create a new ServeMux
	router := http.NewServeMux()

	// Register the user handlers with HTTP methods
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserByIDHandler(w, r)
		case http.MethodPost:
			userHandler.CreateUserHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	// Set up the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}