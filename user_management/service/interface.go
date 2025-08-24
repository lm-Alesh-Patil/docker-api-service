package service

import "context"

// UserServiceInterface defines business logic
type UserServiceInterface interface {
	RegisterUser(ctx context.Context, name, email, password string) (int64, error)
}
