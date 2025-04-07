package container

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type DockerRunner struct{}


func NewDockerRunner() *DockerRunner {
	return &DockerRunner{}
}

func (r *DockerRunner) RunContainer(projectDir, linterImage string) (string, error) {
    timeout := 60 * time.Second
    
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    cmd := exec.CommandContext(ctx,
        "docker", "run", "--rm",
        "--memory", "500M",
        "--cpus", "1.0",
        "--network", "none",
        "--read-only",
        "-v", fmt.Sprintf("%s:/code:ro", projectDir),
        linterImage,
    )

    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    
    switch {
    case ctx.Err() == context.DeadlineExceeded:
        return "", fmt.Errorf("linting timeout after %v", timeout)
    case err != nil:
        if out.Len() == 0 {
            return "", fmt.Errorf("docker execution failed (no output): %w", err)
        }
        return out.String(), nil
    default:
        return out.String(), nil
    }
}