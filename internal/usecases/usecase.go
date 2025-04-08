package usecases

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories"
	"SolutionService/internal/services"
	"SolutionService/internal/services/linting"
)

type (
	ExecuteCodeUsecase interface {
		Handle(req dto.ExecuteRequest) (*dto.MultiExecuteResponse, error)
	}

	GetRuntimesUsecase interface {
		Handle() ([]dto.RuntimeResponse, error)
	}

	Usecase struct {
		ExecuteCodeUsecase ExecuteCodeUsecase
		GetRuntimesUsecase GetRuntimesUsecase
	}
)

func NewUsecase(services *services.Service, linterFactory *linting.LinterFactory, repos *repositories.Repository) *Usecase {
	return &Usecase{
		ExecuteCodeUsecase: NewExecuteCodeUsecase(services, linterFactory, repos),
		GetRuntimesUsecase: NewGetRuntimesUsecase(services),
	}
}
