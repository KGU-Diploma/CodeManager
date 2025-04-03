package usecases

import (
	"CodeManager/internal/services"
	"CodeManager/internal/dto"
)

type (
	ExecuteCodeUsecase interface {
		Handle(req dto.ExecuteRequest) (*dto.ExecuteResponse, []string, error)
	}

	GetRuntimesUsecase interface {
		Handle() ([]dto.RuntimeResponse, error)
	}

	Usecase struct {
		ExecuteCodeUsecase ExecuteCodeUsecase
		GetRuntimesUsecase GetRuntimesUsecase
	}
)

func NewUsecase(services *services.Service) *Usecase {
	return &Usecase{ExecuteCodeUsecase: NewExecuteCodeUsecase(services), GetRuntimesUsecase: NewGetRuntimesUsecase(services)}
}
