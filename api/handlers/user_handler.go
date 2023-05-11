/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:07:12 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eulbyvan/idk/app/go-user-management/internal/app"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
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
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// Implement other user handler methods as needed
