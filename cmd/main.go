package main

import (
	"fmt"
	"log"
	"log/slog"
	"strings"
	"github.com/gin-gonic/gin"

	"CodeManager/internal/pkg/config"
	"CodeManager/internal/pkg/logger"
	"CodeManager/api"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal("Could not create config %v", err)
	}

	logger.InitLogger(cfg.Logger)
	logger := slog.Default()

	router := gin.Default()

	// initialize handler and register routes
	handler := api.NewHandler()
	handler.SetupRoutes(router)

	address := fmt.Sprintf("%s:%s", "127.0.0.1", strings.TrimPrefix(cfg.APP_PORT, ":"))
	logger.Info("Starting server...", "address", address)

	if err := router.Run(address); err != nil {
		logger.Error("Server failed to start: %v", err)
	}
}