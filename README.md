# SolutionService

## How to Run the Service

1. **Clone the repository and navigate to the project directory**
    ```bash
    git clone git@github.com:KGU-Diploma/SolutionService.git
    cd CodeManager
    ```
2. Install dependencies
    ```bash
    go mod download
    ```
3. Navigate to the cmd directory and run the application
    ```bash
    cd cmd
    go run main.go
    ```
4. Check liveness probe
    Use the following command to check if the service is running:
    ```bash
    curl --location 'http://localhost:8001/health/ping'
    ```
