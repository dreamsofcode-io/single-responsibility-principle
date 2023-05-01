package user

import (
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, u User) error {
	_, err := r.db.ExecContext(ctx, saveSQL, u.ID, u.Email, u.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
