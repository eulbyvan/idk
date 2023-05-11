package app

import (
	"errors"

	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

type ProfileService struct {
	profileRepo domain.ProfileRepository
}

func NewProfileService(profileRepo domain.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
	}
}

func (s *ProfileService) CreateProfile(profile *domain.Profile) error {
	// Validate the profile fields
	if profile.UserID == 0 {
		return errors.New("user ID is required")
	}

	// Create the profile in the repository
	err := s.profileRepo.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}

// Implement other profile service methods as needed
