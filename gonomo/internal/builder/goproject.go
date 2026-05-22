package builder

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"gonomo/internal/config"
	runtimeTemplate "gonomo/runtime"
)

func GenerateProject(cfg *config.Config, opts BuildOptions, nodeExe string) (string, error) {
	buildDir := filepath.Join(".gonomo", "build")

	if opts.Clean {
		os.RemoveAll(buildDir)
	}

	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return "", err
	}

	err := fs.WalkDir(runtimeTemplate.TemplatesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dst := filepath.Join(buildDir, path)
		if path == filepath.Join("cmd", "runtime", "main.go") || path == "cmd/runtime/main.go" {
			dst = filepath.Join(buildDir, "main.go")
		}

		if d.IsDir() {
			return os.MkdirAll(dst, 0755)
		}

		content, err := runtimeTemplate.TemplatesFS.ReadFile(path)
		if err != nil {
			return err
		}

		if filepath.Base(path) == "templates.go" {
			return nil
		}

		return os.WriteFile(dst, content, 0644)
	})

	if err != nil {
		return "", fmt.Errorf("extract templates: %w", err)
	}

	// Generate config_values.go with baked-in config
	if err := generateConfigValues(buildDir, cfg); err != nil {
		return "", fmt.Errorf("generate config values: %w", err)
	}

	// Copy icon
	if cfg.Icon != "" {
		iconData, err := os.ReadFile(cfg.Icon)
		if err == nil {
			os.WriteFile(filepath.Join(buildDir, "internal", "resources", "icon.ico"), iconData, 0644)
			os.WriteFile(filepath.Join(buildDir, "icon.ico"), iconData, 0644)
		}
	}

	// Copy Node.exe
	if cfg.Build.Runtime == "node" && nodeExe != "" {
		nodeData, err := os.ReadFile(nodeExe)
		if err != nil {
			return "", fmt.Errorf("read node.exe: %w", err)
		}
		if err := os.WriteFile(filepath.Join(buildDir, "internal", "resources", "node.exe"), nodeData, 0644); err != nil {
			return "", err
		}
	}

	// Copy frontend output
	if cfg.Build.Embed == "full" {
		frontendSrc := filepath.Join(cfg.Build.Cwd, cfg.Build.OutputDir)
		frontendDst := filepath.Join(buildDir, "internal", "resources", "frontend")
		fmt.Printf("Copying frontend from %s to %s...\n", frontendSrc, frontendDst)
		err := copyDir(frontendSrc, frontendDst)
		if err != nil {
			fmt.Printf("Warning: failed to copy frontend: %v\n", err)
		}
	}

	// Generate rsrc.syso
	if cfg.Icon != "" {
		if _, err := os.Stat(filepath.Join(buildDir, "icon.ico")); err == nil {
			fmt.Println("Generating icon resource via rsrc...")
			rsrcCmd := exec.Command("rsrc", "-ico", "icon.ico", "-o", "rsrc.syso")
			rsrcCmd.Dir = buildDir
			if err := rsrcCmd.Run(); err != nil {
				fmt.Printf("Warning: rsrc failed: %v\n", err)
			}
		} else {
			fmt.Println("Warning: icon file not found, skipping rsrc generation.")
		}
	}

	return buildDir, nil
}

func generateConfigValues(buildDir string, cfg *config.Config) error {
	code := fmt.Sprintf(`package config

func init() {
	AppConfig = Config{
		Name:  %q,
		Title: %q,
		Icon:  %q,
		Window: WindowConfig{
			Width:               %d,
			Height:              %d,
			Maximized:           %t,
			TitleBarHidden:      %t,
			TitleBarOverlay:     %t,
			DarkMode:            %t,
			CaptionColor:        %q,
			TextColor:           %q,
			CaptionButtonWidth:  %d,
			CaptionButtonHeight: %d,
		},
		Splash: SplashConfig{
			Enabled:         %t,
			Layout:          %q,
			BackgroundColor: %q,
			ForegroundColor: %q,
			Image:           %q,
			Text:            %q,
			Tagline:         %q,
			MinDurationMs:   %d,
			Width:           %d,
			Height:          %d,
		},
		Build: BuildConfig{
			Command:   %q,
			Cwd:       %q,
			OutputDir: %q,
			Entry:     %q,
			Runtime:   %q,
			Embed:     %q,
		},
		Output: OutputConfig{
			Dir:  %q,
			Name: %q,
		},
	}
}
`,
		cfg.Name, cfg.Title, cfg.Icon,
		cfg.Window.Width, cfg.Window.Height, cfg.Window.Maximized,
		cfg.Window.TitleBarHidden, cfg.Window.TitleBarOverlay, cfg.Window.DarkMode,
		cfg.Window.CaptionColor, cfg.Window.TextColor,
		cfg.Window.CaptionButtonWidth, cfg.Window.CaptionButtonHeight,
		cfg.Splash.Enabled, cfg.Splash.Layout, cfg.Splash.BackgroundColor,
		cfg.Splash.ForegroundColor, cfg.Splash.Image, cfg.Splash.Text,
		cfg.Splash.Tagline, cfg.Splash.MinDurationMs, cfg.Splash.Width, cfg.Splash.Height,
		cfg.Build.Command, cfg.Build.Cwd, cfg.Build.OutputDir,
		cfg.Build.Entry, cfg.Build.Runtime, cfg.Build.Embed,
		cfg.Output.Dir, cfg.Output.Name,
	)

	configDir := filepath.Join(buildDir, "internal", "config")
	os.MkdirAll(configDir, 0755)
	return os.WriteFile(filepath.Join(configDir, "config_values.go"), []byte(code), 0644)
}

func GoBuild(tempDir string, cfg *config.Config) error {
	outName := cfg.Output.Name
	if outName == "" {
		outName = cfg.Name + ".exe"
	}
	outPath, _ := filepath.Abs(filepath.Join(cfg.Output.Dir, outName))

	os.MkdirAll(filepath.Dir(outPath), 0755)

	fmt.Println("Fetching Go dependencies...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = tempDir
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr
	if err := tidyCmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy failed: %w", err)
	}

	cmd := exec.Command("go", "build", "-ldflags", "-s -w -H=windowsgui", "-o", outPath, ".")
	cmd.Dir = tempDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyDir(src string, dst string) error {
	cmd := exec.Command("powershell", "-c", fmt.Sprintf("Copy-Item -Path '%s' -Destination '%s' -Recurse -Force", src, dst))
	return cmd.Run()
}
