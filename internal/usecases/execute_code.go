package usecases

import (
	"CodeManager/internal/dto"
	"CodeManager/internal/services"
	"fmt"
)

type ExecuteCodeUsecaseImpl struct {
	services *services.Service
	linterFactory *services.LinterFactory
}

func NewExecuteCodeUsecase(service *services.Service) ExecuteCodeUsecase {
	return &ExecuteCodeUsecaseImpl{
		services: service,
		linterFactory: services.NewLinterFactory(),
	}
}

func (u *ExecuteCodeUsecaseImpl) Handle(req dto.ExecuteRequest) (*dto.ExecuteResponse, []string, error) {
	result, err := u.services.Piston.ExecuteCode(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute code: %w", err)
	}
	
	linter, err := u.linterFactory.NewLinter(req.Language)
	if err != nil {
		return nil, nil, err
	}

	linterIssues, err := linter.Lint(req.Files[0].Content)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to lint code: %w", err)
	}

	if len(linterIssues) > 0 {
		return result, linterIssues, nil
	}

	return result, nil, nil
}
