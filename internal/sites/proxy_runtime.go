package sites

import (
	"context"
	"errors"
	"strings"
	"time"
)

// A discovered configuration file is insufficient evidence of a working proxy.
// Do not expose raw host diagnostics through a job status message.
func verifyProxyRuntime(ctx context.Context, runner recipeJobCommandRunner) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	output, err := runner.Run(ctx, "docker", "inspect", "--format", "{{.State.Running}}", "nginx")
	if err != nil || strings.TrimSpace(string(output)) != "true" {
		return errors.New("Nginx container is not running")
	}
	if _, err := runner.Run(ctx, "docker", "exec", "nginx", "nginx", "-t"); err != nil {
		return errors.New("Nginx configuration validation failed")
	}
	return nil
}
