/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:07:12 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eulbyvan/idk/app/go-user-management/internal/app"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain/dto/req"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain/dto/res"
	"github.com/eulbyvan/idk/app/go-user-management/pkg"
)

type UserHandler struct {
	userService *app.UserService
}

func NewUserHandler(userService *app.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	log.Println("HEREEEE1", user)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("HEREEEE2", user)

	err = h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request req.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Authenticate the user
	err := h.userService.AuthenticateUser(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := pkg.GenerateToken(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	response := res.LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}

// Implement other user handler methods as needed
