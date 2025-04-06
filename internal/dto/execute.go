package dto

// ExecuteRequest represents the request data for execution.
// @Description Execute a code with specified parameters.
type PistonExecuteRequest struct {
    Language           string        `json:"language" validate:"required" swaggo:"description=The language to use for execution"`
    Version            string        `json:"version" validate:"required" swaggo:"description=The version of the language"`
    Files              []File        `json:"files" validate:"required,dive" swaggo:"description=The files to be executed"`
    Stdin              string        `json:"stdin,omitempty" swaggo:"description=Input data to be passed to stdin"`
    Args               []string      `json:"args,omitempty" swaggo:"description=Arguments to pass to the program"`
    CompileTimeout     int           `json:"compile_timeout,omitempty" default:"10000" swaggo:"description=The compile timeout"`
    RunTimeout         int           `json:"run_timeout,omitempty" default:"3000" swaggo:"description=The run timeout"`
    CompileCPUTime     int           `json:"compile_cpu_time,omitempty" default:"10000" swaggo:"description=The compile CPU time"`
    RunCPUTime         int           `json:"run_cpu_time,omitempty" default:"3000" swaggo:"description=The run CPU time"`
    CompileMemoryLimit int           `json:"compile_memory_limit,omitempty" default:"-1" swaggo:"description=The compile memory limit"`
    RunMemoryLimit     int           `json:"run_memory_limit,omitempty" default:"-1" swaggo:"description=The run memory limit"`
}

// File represents a file to be used in the execution request.
type File struct {
    Name     string `json:"name,omitempty" swaggo:"description=The name of the file"`
    Content  string `json:"content" validate:"required" swaggo:"description=The content of the file"`
    Encoding string `json:"encoding,omitempty" default:"utf8" swaggo:"description=The file content encoding"`
}

// ExecuteResponse represents the response after execution.
type PistonExecuteResponse struct {
    Run Run `json:"run"`  // The result of the execution.
    Language string `json:"language"` // The language used for execution.
    Version string `json:"version"`   // The version of the language used.
}

// Run contains the result of an execution.
type Run struct {
    Signal  string `json:"signal"` // The signal received during execution.
    Stdout  string `json:"stdout"` // The output to stdout.
    Stderr  string `json:"stderr"` // The error output to stderr.
    Code    int    `json:"code"`   // The exit code of the program.
    Output  string `json:"output"` // The final output.
    Memory  int    `json:"memory"` // Memory used by the program.
    Message string `json:"message"` // Message related to the execution result.
    Status  string `json:"status"` // The status of the execution.
    CPUTime int    `json:"cpu_time"` // CPU time spent in execution.
    WallTime int   `json:"wall_time"` // Wall time taken for execution.
}

// TestCaseResult represents the result of a single test case execution.
// @Description Result of a single test case execution.
type TestCaseResult struct {
	Input    string `json:"input" swaggo:"description=The input provided to the program"`         // The input provided to the program
	Expected string `json:"expected" swaggo:"description=The expected output of the program"`      // The expected output of the program
	Actual   string `json:"actual" swaggo:"description=The actual output produced by the program"` // The actual output produced by the program
	Passed   bool   `json:"passed" swaggo:"description=Whether the output matched the expected"`   // Whether the output matched the expected output
	Message  string `json:"message" swaggo:"description=Details about the result or any error"`    // Details about the result or any error
}

// MultiExecuteResponse represents the response after running all test cases.
// @Description Response containing the results of test execution and linting issues.
type MultiExecuteResponse struct {
	Language   string           `json:"language" swaggo:"description=The programming language used"`           // The programming language used
	Version    string           `json:"version" swaggo:"description=The version of the language used"`         // The version of the language used
	Results    []TestCaseResult `json:"results" swaggo:"description=The results of each test case executed"`   // The results of each test case executed
	LintIssues []string         `json:"lint_issues,omitempty" swaggo:"description=List of linting issues"`     // List of linting issues, if any
}


type ExecuteRequest struct {
    TaskId string `json:"task_id" validate:"required" swaggo:"description=The ID of the task"`
    PistonExecuteRequest PistonExecuteRequest `json:"piston_execute_request" validate:"required" swaggo:"description=The request data for execution by piston engine"`
}