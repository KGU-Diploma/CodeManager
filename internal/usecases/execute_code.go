package usecases

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories"
	"SolutionService/internal/services"
	"SolutionService/internal/services/linting"
	"SolutionService/internal/services/tools"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type ExecuteCodeUsecaseImpl struct {
	services      *services.Service
	linterFactory *linting.LinterFactory
	repos         *repositories.Repository
}

func NewExecuteCodeUsecase(service *services.Service, linterFactory *linting.LinterFactory, repo *repositories.Repository) ExecuteCodeUsecase {
	return &ExecuteCodeUsecaseImpl{
		services:      service,
		linterFactory: linterFactory,
		repos:         repo,
	}
}

func (u *ExecuteCodeUsecaseImpl) Handle(req dto.ExecuteRequest) (*dto.MultiExecuteResponse, error) {
	testData, err := u.repos.TestData.GetTestDataByTaskId("f75a267e-0756-49fb-984b-82f9e2b5a5fb")
	if err != nil {
		slog.Error("Error getting test data by task ID", "error", err)
		return nil, fmt.Errorf("failed to get test data: %w", err)
	}

	var testResults []dto.TestCaseResult

	for _, test := range testData {
		req.PistonExecuteRequest.Stdin = test.Input

		pistonResult, err := u.services.Piston.ExecuteCode(req.PistonExecuteRequest)
		if err != nil {
			testResults = append(testResults, dto.TestCaseResult{
				Input:    test.Input,
				Expected: test.Output,
				Actual:   "",
				Passed:   false,
				Message:  fmt.Sprintf("execution error: %v", err),
			})
			continue
		}

		passed := tools.CompareExpectedAndActual(test.Output, pistonResult.Run.Stdout)

		testResults = append(testResults, dto.TestCaseResult{
			Input:    test.Input,
			Expected: test.Output,
			Actual:   pistonResult.Run.Stdout,
			Passed:   passed,
			Message:  "",
		})
	}

	codeLanguage := req.PistonExecuteRequest.Language
	linter, err := u.linterFactory.NewLinter(codeLanguage)
	if err != nil {
		slog.Error("Error creating linter", "error", err)
		return nil, fmt.Errorf("failed to create linter: %w", err)
	}

	fileContent := req.PistonExecuteRequest.Files[0].Content

	lintIssues, err := linter.Lint(fileContent)
	if err != nil {
		slog.Error("Error creating linter", "error", err)
		return nil, fmt.Errorf("failed to lint code: %w", err)
	}
	
	err = u.services.SolutionService.CreateCodingSolution(req.TaskId, uuid.UUID{}, testResults, fileContent, codeLanguage, lintIssues)
	response := &dto.MultiExecuteResponse{
		Language:   codeLanguage,
		Version:    req.PistonExecuteRequest.Version,
		Results:    testResults,
		LintIssues: lintIssues,
	}

	return response, nil
}
