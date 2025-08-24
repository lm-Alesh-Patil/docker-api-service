package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"

	"github.com/lm-Alesh-Patil/docker-api-service/config"
	"github.com/lm-Alesh-Patil/docker-api-service/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load config from YAML
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// MySQL connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DB.MySQL.Username,
		cfg.DB.MySQL.Password,
		cfg.DB.MySQL.Host,
		cfg.DB.MySQL.Port,
		cfg.DB.MySQL.Database,
	))
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DB.Redis.Host, cfg.DB.Redis.Port),
		Password: cfg.DB.Redis.Password,
		DB:       cfg.DB.Redis.DB,
	})
	defer rdb.Close()

	// Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Register application routes
	routes.RegisterRoutes(router, *cfg, db, rdb)

	// HTTP server
	addr := fmt.Sprintf("%s:%d", cfg.Connection.HTTP.Host, cfg.Connection.HTTP.Port)
	fmt.Printf("Server running at http://%s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
