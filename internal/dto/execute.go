package dto

type ExecuteRequest struct {
    Language           string        `json:"language" validate:"required"`                      // The language to use for execution, must be a string and must be installed.
    Version            string        `json:"version" validate:"required"`                       // The version of the language to use for execution, must be a string containing a SemVer selector or specific version number.
    Files              []File        `json:"files" validate:"required,dive"`                    // An array of files containing code or other data for execution. The first file is considered the main file.
    Stdin              string        `json:"stdin,omitempty"`                                   // The text to pass as stdin to the program. Defaults to blank string.
    Args               []string      `json:"args,omitempty"`                                    // The arguments to pass to the program. Defaults to [].
    CompileTimeout     int           `json:"compile_timeout,omitempty" default:"10000"`         // The maximum wall-time for the compile stage in milliseconds. Defaults to 10000 (10 seconds).
    RunTimeout         int           `json:"run_timeout,omitempty" default:"3000"`              // The maximum wall-time for the run stage in milliseconds. Defaults to 3000 (3 seconds).
    CompileCPUTime     int           `json:"compile_cpu_time,omitempty" default:"10000"`        // The maximum CPU-time for the compile stage in milliseconds. Defaults to 10000 (10 seconds).
    RunCPUTime         int           `json:"run_cpu_time,omitempty" default:"3000"`             // The maximum CPU-time for the run stage in milliseconds. Defaults to 3000 (3 seconds).
    CompileMemoryLimit int           `json:"compile_memory_limit,omitempty" default:"-1"`       // The maximum memory for the compile stage in bytes. Defaults to -1 (no limit).
    RunMemoryLimit     int           `json:"run_memory_limit,omitempty" default:"-1"`           // The maximum memory for the run stage in bytes. Defaults to -1 (no limit).
}

// File represents a file to be used in the execution request.
type File struct {
    Name     string `json:"name,omitempty"`                    // The name of the file to upload, must be a string containing no path or left out.
    Content  string `json:"content" validate:"required"`       // The content of the file to upload, must be a string containing text to write.
    Encoding string `json:"encoding,omitempty" default:"utf8"` // The encoding scheme used for the file content. Defaults to utf8.
}

type ExecuteResponse struct {
	Run Run `json:"run"` 			  // The result of the execution.
	Language string `json:"language"` // The language used for execution.
	Version string `json:"version"`   // The version of the language used for execution.
}

type Run struct {
	Signal string `json:"signal"` //
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Code	int    `json:"code"`
	Output string `json:"output"`
	Memory int `json:"memory"`
	Message string `json:"message"`
	Status string `json:"status"`
	CPUTime int `json:"cpu_time"`
	WallTime int `json:"wall_time"`
}