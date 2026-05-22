package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type WindowConfig struct {
	Width     int             `json:"width"`
	Height    int             `json:"height"`
	Maximized bool            `json:"maximized,omitempty"`
	RawStyle  json.RawMessage `json:"titleBarStyle,omitempty"`

	// Parsed fields
	TitleBarHidden      bool   `json:"-"`
	TitleBarOverlay     bool   `json:"-"`
	DarkMode            bool   `json:"-"`
	CaptionColor        string `json:"-"`
	TextColor           string `json:"-"`
	CaptionButtonWidth  int    `json:"-"`
	CaptionButtonHeight int    `json:"-"`
}

func (w *WindowConfig) parseTitleBarStyle() {
	if len(w.RawStyle) == 0 {
		return
	}
	var s string
	if err := json.Unmarshal(w.RawStyle, &s); err == nil {
		if s == "hidden" {
			w.TitleBarHidden = true
		}
		return
	}

	var obj struct {
		Hidden              bool   `json:"hidden"`
		Overlay             bool   `json:"overlay"`
		DarkMode            bool   `json:"darkMode"`
		CaptionColor        string `json:"captionColor"`
		TextColor           string `json:"textColor"`
		CaptionButtonWidth  int    `json:"captionButtonWidth"`
		CaptionButtonHeight int    `json:"captionButtonHeight"`
	}
	if err := json.Unmarshal(w.RawStyle, &obj); err == nil {
		w.TitleBarHidden = obj.Hidden
		w.TitleBarOverlay = obj.Overlay
		w.DarkMode = obj.DarkMode
		w.CaptionColor = obj.CaptionColor
		w.TextColor = obj.TextColor
		w.CaptionButtonWidth = obj.CaptionButtonWidth
		w.CaptionButtonHeight = obj.CaptionButtonHeight
	}
}

type SplashConfig struct {
	Enabled         bool   `json:"enabled"`
	Layout          string `json:"layout"`
	BackgroundColor string `json:"backgroundColor"`
	ForegroundColor string `json:"foregroundColor"`
	Image           string `json:"image"`
	Text            string `json:"text"`
	Tagline         string `json:"tagline,omitempty"`
	MinDurationMs   int    `json:"minDuration"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
}

type BuildConfig struct {
	Command     string `json:"command"`
	Cwd         string `json:"cwd"`
	OutputDir   string `json:"outputDir"`
	Entry       string `json:"entry"`
	Runtime     string `json:"runtime"`
	Embed       string `json:"embed"`
	NodeVersion string `json:"nodeVersion,omitempty"`
}

type DevConfig struct {
	Command string `json:"command"`
	Cwd     string `json:"cwd"`
	Url     string `json:"url"`
}

type OutputConfig struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
}

type Config struct {
	Name   string       `json:"name"`
	Title  string       `json:"title"`
	Icon   string       `json:"icon"`
	Window WindowConfig `json:"window"`
	Splash SplashConfig `json:"splash"`
	Build  BuildConfig  `json:"build"`
	Dev    DevConfig    `json:"dev"`
	Output OutputConfig `json:"output"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	cfg.Window.parseTitleBarStyle()
	return &cfg, nil
}

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
