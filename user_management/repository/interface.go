package repository

import (
	"context"

	"github.com/lm-Alesh-Patil/docker-api-service/user_management/models"
)

// UserRepositoryInterface defines DB operations
type UserRepositoryInterface interface {
	SaveUser(ctx context.Context, user models.User) (int64, error)
}
