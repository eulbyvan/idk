/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:01:49 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) CreateProfile(profile *domain.Profile) error {
	query := `
		INSERT INTO profiles (user_id, bio)
		VALUES ($1, $2)
		RETURNING id
	`
	err := r.db.QueryRow(query, profile.UserID, profile.Bio).Scan(&profile.ID)
	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}
	return nil
}

// Implement other profile repository methods as needed (e.g., GetProfileByUserID, UpdateProfile, DeleteProfile)
