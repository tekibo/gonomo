package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"gonomo/internal/config"
)

type nodeRelease struct {
	Version string      `json:"version"`
	Lts     interface{} `json:"lts"`
}

func fetchLatestLTSVersion() (string, error) {
	fmt.Println("Fetching latest Node.js LTS version...")
	resp, err := http.Get("https://nodejs.org/dist/index.json")
	if err != nil {
		return "", fmt.Errorf("failed to fetch node releases: %w", err)
	}
	defer resp.Body.Close()

	var releases []nodeRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", fmt.Errorf("failed to decode node releases: %w", err)
	}

	for _, rel := range releases {
		// lts is false or a string
		if ltsStr, ok := rel.Lts.(string); ok && ltsStr != "" {
			return rel.Version, nil
		}
	}

	return "", fmt.Errorf("could not find any LTS release")
}

func EnsureRuntimeBin(cfg *config.Config) (string, error) {
	if cfg.Build.Runtime != "node" {
		return "", fmt.Errorf("runtime '%s' is not supported yet (only 'node' or 'static' are currently fully supported)", cfg.Build.Runtime)
	}

	nodeVersion := cfg.Build.NodeVersion
	if nodeVersion == "" {
		var err error
		nodeVersion, err = fetchLatestLTSVersion()
		if err != nil {
			return "", err
		}
	}

	// Make sure version has 'v' prefix
	if nodeVersion[0] != 'v' {
		nodeVersion = "v" + nodeVersion
	}

	binDir := filepath.Join(".gonomo", "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		return "", err
	}

	nodeExe := filepath.Join(binDir, fmt.Sprintf("node-%s.exe", nodeVersion))
	if _, err := os.Stat(nodeExe); err == nil {
		fmt.Printf("Node.js %s already cached.\n", nodeVersion)
		return nodeExe, nil
	}

	url := fmt.Sprintf("https://nodejs.org/dist/%s/win-x64/node.exe", nodeVersion)
	fmt.Printf("Downloading Node.js %s from %s...\n", nodeVersion, url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download node.exe: status %d", resp.StatusCode)
	}

	out, err := os.Create(nodeExe)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Node.js binary downloaded successfully.")
	return nodeExe, nil
}
