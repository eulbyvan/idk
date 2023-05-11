/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:34:01 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/eulbyvan/idk/app/go-user-management/api/handlers"
	"github.com/eulbyvan/idk/app/go-user-management/api/middlewares"
	"github.com/eulbyvan/idk/app/go-user-management/internal/app"
	"github.com/eulbyvan/idk/app/go-user-management/internal/repository"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "postgres", "user_management")) // Initialize your PostgreSQL database connection here

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("failed to run the application")
		} else {
			db.Close()
		}
	}()

	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	userService := app.NewUserService(userRepo, profileRepo)
	profileService := app.NewProfileService(profileRepo)

	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)

	// Middlewares
	r.Use(middlewares.Logger())
	r.Use(middlewares.Authentication()) // Add the authentication middleware

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		// Add other user routes as needed
	}

	// Profile routes
	profileRoutes := r.Group("/profiles")
	{
		profileRoutes.POST("/", profileHandler.CreateProfile)
		// Add other profile routes as needed
	}

	return r
}
