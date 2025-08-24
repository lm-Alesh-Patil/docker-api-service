package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/lm-Alesh-Patil/docker-api-service/config"
	"github.com/lm-Alesh-Patil/docker-api-service/user_management/handler"
	"github.com/lm-Alesh-Patil/docker-api-service/user_management/repository"
	"github.com/lm-Alesh-Patil/docker-api-service/user_management/service"
)

func RegisterRoutes(router *chi.Mux, cfg config.Config, db *sql.DB, redisClient *redis.Client) {
	userRepo := repository.NewMysqlUserRepository(db)
	userService := service.NewUserService(userRepo, redisClient)
	userHandler := handler.NewUserHandler(userService)
	router.Post("/register", userHandler.Register)

}
