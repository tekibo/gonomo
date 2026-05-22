package builder

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"gonomo/internal/config"
)

type BuildOptions struct {
	SkipFrontend bool
	Verbose      bool
	Clean        bool
}

func Build(opts BuildOptions) error {
	cfg, err := config.Load("gonomo.json")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// 1. Run frontend build
	if !opts.SkipFrontend && cfg.Build.Command != "" {
		fmt.Println("Building frontend...")
		if err := RunFrontendBuild(cfg); err != nil {
			return fmt.Errorf("frontend build failed: %w", err)
		}
	}

	// 2. Download/Bundle Runtime Binaries
	var nodeExe string
	if cfg.Build.Runtime == "node" || cfg.Build.Runtime == "bun" {
		fmt.Printf("Ensuring %s runtime is available...\n", cfg.Build.Runtime)
		nodeExe, err = EnsureRuntimeBin(cfg)
		if err != nil {
			return fmt.Errorf("failed to get runtime binary: %w", err)
		}
	}

	// 3. Generate Temp Go Project
	fmt.Println("Generating temporary Go project...")
	tempDir, err := GenerateProject(cfg, opts, nodeExe)
	if err != nil {
		return fmt.Errorf("failed to generate go project: %w", err)
	}

	// 4. Go Build
	fmt.Println("Compiling final executable...")
	if err := GoBuild(tempDir, cfg); err != nil {
		return fmt.Errorf("go build failed: %w", err)
	}

	return nil
}

func waitForURL(url string) error {
	for i := 0; i < 60; i++ {
		conn, err := net.DialTimeout("tcp", url, time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("dev server at %s did not start within 60 seconds", url)
}

func Dev() error {
	cfg, err := config.Load("gonomo.json")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// 1. Start frontend dev server
	fmt.Printf("Starting dev server: %s in %s\n", cfg.Dev.Command, cfg.Dev.Cwd)
	var devCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		devCmd = exec.Command("cmd", "/c", cfg.Dev.Command)
	} else {
		devCmd = exec.Command("sh", "-c", cfg.Dev.Command)
	}
	devCmd.Dir = cfg.Dev.Cwd
	devCmd.Stdout = os.Stdout
	devCmd.Stderr = os.Stderr
	if err := devCmd.Start(); err != nil {
		return fmt.Errorf("failed to start dev server: %w", err)
	}

	// 2. Wait for dev server to be ready
	devURL := cfg.Dev.Url
	// Parse host:port from URL for health check
	hostPort := devURL
	if len(devURL) > 7 && devURL[:7] == "http://" {
		hostPort = devURL[7:]
	} else if len(devURL) > 8 && devURL[:8] == "https://" {
		hostPort = devURL[8:]
	}
	fmt.Printf("Waiting for dev server at %s...\n", devURL)
	if err := waitForURL(hostPort); err != nil {
		devCmd.Process.Kill()
		return err
	}
	fmt.Println("Dev server is ready.")

	// 3. Generate temp Go project (without frontend embed)
	embedBackup := cfg.Build.Embed
	cfg.Build.Embed = "none"
	tempDir, err := GenerateProject(cfg, BuildOptions{SkipFrontend: true}, "")
	if err != nil {
		devCmd.Process.Kill()
		return fmt.Errorf("failed to generate go project: %w", err)
	}
	cfg.Build.Embed = embedBackup

	// 4. Build the Go app
	fmt.Println("Compiling dev executable...")
	cmd := exec.Command("go", "build", "-ldflags", "-s -w", "-o", "gonomo-dev.exe", ".")
	cmd.Dir = tempDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		devCmd.Process.Kill()
		return fmt.Errorf("go build failed: %w", err)
	}

	// 5. Run the dev executable
	devExe, _ := filepath.Abs(filepath.Join(tempDir, "gonomo-dev.exe"))
	fmt.Println("Starting app in dev mode...")
	appCmd := exec.Command(devExe)
	appCmd.Stdout = os.Stdout
	appCmd.Stderr = os.Stderr
	appCmd.Env = append(os.Environ(), fmt.Sprintf("GONOMO_DEV_URL=%s", devURL))
	if err := appCmd.Start(); err != nil {
		devCmd.Process.Kill()
		return fmt.Errorf("failed to start dev app: %w", err)
	}

	// 6. Wait for interrupt and clean up
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nShutting down dev mode...")
	appCmd.Process.Kill()
	devCmd.Process.Kill()
	return nil
}
