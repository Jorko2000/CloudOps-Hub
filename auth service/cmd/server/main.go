package main

import (
	"log"

	"github.com/georgigeorgiev/cloudops/auth-service/internal/config"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/handler"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/repository"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/router"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/service"

	"github.com/georgigeorgiev/cloudops/auth-service/pkg/database"
	"github.com/georgigeorgiev/cloudops/auth-service/pkg/logger"
)

func main() {

	cfg := config.Load()

	err := logger.Init()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repository.PostgresUserRepository{
		DB: db,
	}

	authService := &service.AuthService{
		Repo:      userRepo,
		JWTSecret: cfg.JWTSecret,
	}

	authHandler := &handler.AuthHandler{
		Service: authService,
	}

	r := router.Setup(
		authHandler,
		cfg.JWTSecret,
	)

	logger.Log.Info(
		"Auth Service Started",
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
