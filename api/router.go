/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:34:01 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/eulbyvan/idk/app/go-user-management/api/handlers"
	"github.com/eulbyvan/idk/app/go-user-management/api/middlewares"
	"github.com/eulbyvan/idk/app/go-user-management/internal/app"
	"github.com/eulbyvan/idk/app/go-user-management/internal/repository"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	userService := app.NewUserService(userRepo, profileRepo)
	profileService := app.NewProfileService(profileRepo)

	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)

	// Middlewares
	r.Use(middlewares.Logger())
	r.Use(middlewares.Authentication()) // Add the authentication middleware

	apiV1 := r.Group("/api/v1")

	// User routes
	userRoutes := apiV1.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		// Add other user routes as needed
	}

	// Profile routes
	profileRoutes := apiV1.Group("/profiles")
	{
		profileRoutes.POST("/", profileHandler.CreateProfile)
		// Add other profile routes as needed
	}

	return r
}
