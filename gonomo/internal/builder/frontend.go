package builder

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"gonomo/internal/config"
)

func RunFrontendBuild(cfg *config.Config) error {
	if cfg.Build.Command == "" {
		return nil
	}

	fmt.Printf("Running frontend build: %s in %s\n", cfg.Build.Command, cfg.Build.Cwd)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", cfg.Build.Command)
	} else {
		cmd = exec.Command("sh", "-c", cfg.Build.Command)
	}

	cmd.Dir = cfg.Build.Cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("frontend build command failed: %w", err)
	}

	return nil
}
