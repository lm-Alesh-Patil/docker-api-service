package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lm-Alesh-Patil/docker-api-service/user_management/models"
)

type MysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) UserRepositoryInterface {
	return &MysqlUserRepository{db: db}
}

func (r *MysqlUserRepository) SaveUser(ctx context.Context, user models.User) (int64, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}
