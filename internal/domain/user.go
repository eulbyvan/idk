/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 11 2023 9:59:05 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package domain

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRepository interface {
	GetUserByID(userID string) (*User, error)
	CreateUser(user *User) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserByID(userID string) (*User, error) {
	return s.userRepository.GetUserByID(userID)
}

func (s *UserService) CreateUser(user *User) error {
	// Add any validation or business logic here before creating the user
	return s.userRepository.CreateUser(user)
}