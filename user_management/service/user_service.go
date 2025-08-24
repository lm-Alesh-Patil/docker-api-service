package service

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/lm-Alesh-Patil/docker-api-service/user_management/models"
	"github.com/lm-Alesh-Patil/docker-api-service/user_management/repository"
)

type UserService struct {
	repo  repository.UserRepositoryInterface
	redis *redis.Client
}

func NewUserService(repo repository.UserRepositoryInterface, redis *redis.Client) UserServiceInterface {
	return &UserService{repo: repo, redis: redis}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email, password string) (int64, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	// Save in MySQL
	id, err := s.repo.SaveUser(ctx, user)
	if err != nil {
		return 0, err
	}

	// Push to Redis queue
	msg := fmt.Sprintf("%d|%s|%s", id, name, email)
	if err := s.redis.LPush(ctx, "notification_queue", msg).Err(); err != nil {
		return id, fmt.Errorf("failed to push to redis: %w", err)
	}

	return id, nil
}
