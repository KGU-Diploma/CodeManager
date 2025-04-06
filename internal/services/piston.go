package services

import (
	"CodeManager/internal/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PistonService struct{}

func NewPistonService() *PistonService {
	return &PistonService{}
}

func (ps *PistonService) ExecuteCode(req dto.ExecuteRequest) (*dto.PistonExecuteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	} 

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("http://192.168.0.2:2000/api/v2/execute", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var result dto.PistonExecuteResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (ps *PistonService) GetRuntimes() ([]dto.RuntimeResponse, error){
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://192.168.0.2:2000/api/v2/runtimes")
	if err != nil {
		return nil, fmt.Errorf("failed to send request %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	var result []dto.RuntimeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response %w", err)
	}

	return result, nil
}