/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:31:02 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eulbyvan/idk/app/go-user-management/internal/app"
	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

type ProfileHandler struct {
	profileService *app.ProfileService
}

func NewProfileHandler(profileService *app.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.profileService.CreateProfile(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

// Implement other profile handler methods as needed
