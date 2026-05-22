package app

import (
	"fmt"
	"os"
	"unsafe"

	webview "github.com/webview/webview_go"

	"gonomo/runtime/internal/bridge"
	"gonomo/runtime/internal/config"
	"gonomo/runtime/internal/server"
	"gonomo/runtime/internal/window"
)

func Run(cfg *config.Config) error {
	url := os.Getenv("GONOMO_DEV_URL")
	if url == "" {
		nitro, err := server.Start(cfg)
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}
		defer nitro.Cmd.Process.Kill()
		url = nitro.URL
	}

	splash, err := window.CreateSplash(cfg.Splash)
	if err != nil {
		panic(err)
	}

	os.Setenv("WEBVIEW2_ADDITIONAL_BROWSER_ARGUMENTS",
		"--enable-features=msWebView2EnableDraggableRegions")

	host, err := window.CreateHostWindow(cfg.WindowWidth(), cfg.WindowHeight())
	if err != nil {
		panic(err)
	}

	w := webview.NewWindow(false, unsafe.Pointer(&host))
	defer w.Destroy()

	w.SetTitle(cfg.TitleOrDefault())
	w.SetSize(cfg.WindowWidth(), cfg.WindowHeight(), webview.HintNone)

	win := window.Init(w, cfg.Icon, cfg.Window.Maximized)

	if cfg.Window.CaptionButtonWidth > 0 || cfg.Window.CaptionButtonHeight > 0 {
		window.SetCaptionButtonSizes(int32(cfg.Window.CaptionButtonWidth), int32(cfg.Window.CaptionButtonHeight))
	}

	if cfg.UsesTitleBarOverlay() {
		win.SetTitleBarOverlay(true)
	} else if cfg.ShouldHideTitlebar() {
		win.SetTitlebarVisible(false)
	}

	bridge.BindAll(w, cfg, win, splash)

	w.Navigate(url)
	w.Run()
	return nil
}
