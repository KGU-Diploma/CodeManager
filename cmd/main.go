package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"SolutionService/api"
	"SolutionService/internal/pkg/config"
	"SolutionService/internal/pkg/logger"
	"SolutionService/internal/repositories"
	"SolutionService/internal/services"
	"SolutionService/internal/services/linting"
	"SolutionService/internal/usecases"
	"SolutionService/internal/services/container"
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