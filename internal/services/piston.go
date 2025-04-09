package services

import (
	"SolutionService/internal/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type PistonServiceImpl struct{}

func NewPistonService() *PistonServiceImpl {
	return &PistonServiceImpl{}
}

func (ps *PistonServiceImpl) ExecuteCode(req dto.PistonExecuteRequest) (*dto.PistonExecuteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		slog.Error("Code execution error", "error", err)
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("http://192.168.0.2:2000/api/v2/execute", "application/json", bytes.NewBuffer(body))
	if err != nil {
		slog.Error("Error sending request to piston", "error", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Unexpected status code", "status code", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var result dto.PistonExecuteResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		slog.Error("Error decoding response", "error", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (ps *PistonServiceImpl) GetRuntimes() ([]dto.RuntimeResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://192.168.0.2:2000/api/v2/runtimes")
	if err != nil {
		slog.Error("Error decoding response", "error", err)
		return nil, fmt.Errorf("failed to send request %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	var result []dto.RuntimeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		slog.Error("Error create decoder", "error", err)
		return nil, fmt.Errorf("failed to decode response %w", err)
	}

	return result, nil
}
