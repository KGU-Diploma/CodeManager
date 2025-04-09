package usecases

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/services"
	"fmt"
	"log/slog"
)

type GetRuntimesUsecaseImpl struct {
	services *services.Service
}

func NewGetRuntimesUsecase(services *services.Service) GetRuntimesUsecase {
	return &GetRuntimesUsecaseImpl{services: services}
}

func (u *GetRuntimesUsecaseImpl) Handle() ([]dto.RuntimeResponse, error) {
	runtimes, err := u.services.Piston.GetRuntimes()
	if err != nil {
		slog.Error("Error getting runtimes", "error", err)
		return nil, fmt.Errorf("failed to get runtimes: %w", err)
	}
	return runtimes, nil
}
