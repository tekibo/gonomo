package config

type WindowConfig struct {
	Width               int
	Height              int
	Maximized           bool
	TitleBarHidden      bool
	TitleBarOverlay     bool
	DarkMode            bool
	CaptionColor        string
	TextColor           string
	CaptionButtonWidth  int
	CaptionButtonHeight int
}

type SplashConfig struct {
	Enabled         bool
	Layout          string
	BackgroundColor string
	ForegroundColor string
	Image           string
	Text            string
	Tagline         string
	MinDurationMs   int
	Width           int
	Height          int
}

type BuildConfig struct {
	Entry     string
	Runtime   string
	Embed     string
	Command   string
	Cwd       string
	OutputDir string
}

type DevConfig struct {
	Command string
	Cwd     string
	Url     string
}

type OutputConfig struct {
	Dir  string
	Name string
}

type Config struct {
	Name   string
	Title  string
	Icon   string
	Window WindowConfig
	Splash SplashConfig
	Build  BuildConfig
	Dev    DevConfig
	Output OutputConfig
}

var AppConfig Config

func (c *Config) TitleOrDefault() string {
	if c.Title != "" {
		return c.Title
	}
	if c.Name != "" {
		return c.Name
	}
	return "App"
}

func (c *Config) WindowWidth() int {
	if c.Window.Width > 0 {
		return c.Window.Width
	}
	return 1400
}

func (c *Config) WindowHeight() int {
	if c.Window.Height > 0 {
		return c.Window.Height
	}
	return 900
}

func (c *Config) UsesTitleBarOverlay() bool {
	return c.Window.TitleBarOverlay
}

func (c *Config) ShouldHideTitlebar() bool {
	if c.Window.TitleBarHidden {
		return true
	}
	return false
}

func (s *SplashConfig) LayoutOrDefault() string {
	switch s.Layout {
	case "centered", "minimal", "top-banner", "bottom-banner", "split", "full-image", "custom":
		return s.Layout
	}
	return "centered"
}

func (s *SplashConfig) WidthOrDefault() int {
	if s.Width > 0 {
		return s.Width
	}
	return 480
}

func (s *SplashConfig) HeightOrDefault() int {
	if s.Height > 0 {
		return s.Height
	}
	return 320
}

func (s *SplashConfig) MinDurationOrDefault() int {
	if s.MinDurationMs > 0 {
		return s.MinDurationMs
	}
	return 400
}
