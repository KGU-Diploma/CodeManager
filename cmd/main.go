package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"CodeManager/api"
	"CodeManager/internal/pkg/config"
	"CodeManager/internal/pkg/logger"
	"CodeManager/internal/repositories"
	"CodeManager/internal/services"
	"CodeManager/internal/services/linting"
	"CodeManager/internal/usecases"
	"CodeManager/internal/services/container"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal("Could not create config %v", err)
	}

	logger.InitLogger(cfg.Logger)
	logger := slog.Default()

	db, err := repositories.NewPostgresConnection(cfg.DB_CONNECTION_STRING)
	if err != nil {
		// todo slog.Fatal("Could not connect to database %v", err)
	}
	repos := repositories.NewRepository(db)
	service := services.NewService()
	runner := container.NewDockerRunner()
	linterFactory := linting.NewLinterFactory(runner)
	usecases := usecases.NewUsecase(service, linterFactory, repos)
	handler := api.NewHandler(usecases)
	gitEngine := handler.SetupRoutes()

	httpServer := &http.Server{
		Addr:           ":" + cfg.APP_PORT,
		Handler:        gitEngine,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error("Error starting server", "error", err)
	}

}