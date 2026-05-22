package server

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/phayes/freeport"
	"gonomo/runtime/internal/config"
	"gonomo/runtime/internal/resources"
)

type Server struct {
	Port int
	Cmd  *exec.Cmd
	URL  string
}

func waitForServer(port int) error {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	for i := 0; i < 50; i++ {
		conn, err := net.DialTimeout("tcp", addr, time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}
	return fmt.Errorf("server failed to start or did not bind to port %d", port)
}

func Start(cfg *config.Config) (*Server, error) {
	// 1. Extract embedded resources to a temp folder
	resDir, err := resources.ExtractAll(cfg.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to extract resources: %w", err)
	}

	port, err := freeport.GetFreePort()
	if err != nil {
		return nil, err
	}

	var cmd *exec.Cmd

	if cfg.Build.Runtime == "node" {
		nodeExe := filepath.Join(resDir, "node.exe")
		entryPath := filepath.Join(resDir, "frontend", cfg.Build.Entry)

		cmd = exec.Command(nodeExe, entryPath)
		cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))
	} else if cfg.Build.Runtime == "static" {
		// Start Go static server (will implement separately in static.go.tmpl if needed)
		// For now just error if it reaches here and we don't have static logic yet.
		return nil, fmt.Errorf("static runtime not implemented in this server file")
	} else {
		return nil, fmt.Errorf("unknown runtime: %s", cfg.Build.Runtime)
	}

	// CREATE_NO_WINDOW prevents the child node process from spawning
	// its own visible console window on Windows.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if err := waitForServer(port); err != nil {
		cmd.Process.Kill()
		return nil, err
	}

	return &Server{
		Port: port,
		Cmd:  cmd,
		URL:  fmt.Sprintf("http://127.0.0.1:%d", port),
	}, nil
}
