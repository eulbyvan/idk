/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:12:39 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package app

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

const (
	minPasswordLength = 8
)

type UserService struct {
	userRepo    domain.UserRepository
	profileRepo domain.ProfileRepository
}

func NewUserService(userRepo domain.UserRepository, profileRepo domain.ProfileRepository) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	// Validate the user fields
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < minPasswordLength {
		return errors.New("password must be at least 8 characters long")
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user in the repository
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	// Create a profile for the user
	profile := &domain.Profile{
		UserID: user.ID,
		Bio:    "",
	}
	err = s.profileRepo.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) AuthenticateUser(email string, password string) error {
	// Call the repository method to get the user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	if password != user.Password {
		return errors.New("invalid credentials")
	}

	return nil
}
// Implement other user service methods as needed
