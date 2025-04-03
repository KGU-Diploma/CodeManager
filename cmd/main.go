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


	// initialize handler and register routes
	services, err := services.NewService("python")
	usecases := usecases.NewUsecase(services)
	handler := api.NewHandler(usecases)
	gitEngine := handler.SetupRoutes()

	// address := fmt.Sprintf("%s:%s", "127.0.0.1", strings.TrimPrefix(cfg.APP_PORT, ":"))
	// logger.Info("Starting server...", "address", address)

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