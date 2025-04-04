package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"CodeManager/api"
	"CodeManager/internal/pkg/config"
	"CodeManager/internal/pkg/logger"
	"CodeManager/internal/services"
	"CodeManager/internal/usecases"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal("Could not create config %v", err)
	}

	logger.InitLogger(cfg.Logger)
	logger := slog.Default()

	service := services.NewService()
	usecases := usecases.NewUsecase(service)
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