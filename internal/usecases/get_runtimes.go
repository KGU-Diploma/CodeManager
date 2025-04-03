package usecases

import (
	"CodeManager/internal/dto"
	"CodeManager/internal/services"
	"fmt"
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
		return nil, fmt.Errorf("failed to get runtimes: %w", err)
	}
	return runtimes, nil
}
