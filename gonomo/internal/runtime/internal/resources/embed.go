package resources

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed all:*
var Files embed.FS

const IconFile = "icon.ico"

func Extract(name string) (string, error) {
	data, err := Files.ReadFile(name)
	if err != nil {
		return "", err
	}
	dir := filepath.Join(os.TempDir(), "gonomo-resources")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	dst := filepath.Join(dir, name)
	if err := os.WriteFile(dst, data, 0644); err != nil {
		return "", err
	}
	return dst, nil
}

// ExtractAll extracts all embedded resources to a temporary directory
// and returns the path to that directory.
func ExtractAll(appName string) (string, error) {
	// Use a predictable directory in Temp based on app name to avoid 
	// re-extracting if not necessary. In a real app, you might want to 
	// hash the app version to know when to overwrite.
	dir := filepath.Join(os.TempDir(), fmt.Sprintf("%s-resources", appName))
	
	// Create the root directory
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	err := fs.WalkDir(Files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		dst := filepath.Join(dir, path)
		
		if d.IsDir() {
			return os.MkdirAll(dst, 0755)
		}

		// Read from embedded FS
		data, err := Files.ReadFile(path)
		if err != nil {
			return err
		}

		// Write to disk
		return os.WriteFile(dst, data, 0644)
	})

	if err != nil {
		return "", err
	}

	return dir, nil
}
