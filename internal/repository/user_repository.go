package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/eulbyvan/idk/app/go-user-management/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// Implement other user repository methods as needed (e.g., GetUserByID, UpdateUser, DeleteUser)
