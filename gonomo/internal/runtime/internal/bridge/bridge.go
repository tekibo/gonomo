package bridge

import (
	"encoding/json"
	"fmt"
	"strconv"

	webview "github.com/webview/webview_go"

	"gonomo/runtime/internal/config"
	"gonomo/runtime/internal/window"
)

func BindAll(w webview.WebView, cfg *config.Config, win *window.Window, splash *window.SplashWindow) {
	cfgJSON, _ := json.Marshal(cfg)

	// Apply titlebar config immediately (before frontend loads)
	win.SetDarkMode(cfg.Window.DarkMode)
	if cfg.Window.CaptionColor != "" {
		win.SetCaptionColor(cfg.Window.CaptionColor)
	}
	if cfg.Window.TextColor != "" {
		win.SetTextColor(cfg.Window.TextColor)
	}

	w.Bind("gonomoClose", func() { win.Close() })
	w.Bind("gonomoMinimize", func() { win.Minimize() })
	w.Bind("gonomoMaximize", func() { win.Maximize() })
	w.Bind("gonomoRestore", func() { win.Restore() })
	w.Bind("gonomoIsMaximized", func() bool { return win.IsMaximized() })
	w.Bind("gonomoSetDarkMode", func(enabled bool) { win.SetDarkMode(enabled) })
	w.Bind("gonomoSetCaptionColor", func(hex string) { win.SetCaptionColor(hex) })
	w.Bind("gonomoSetTextColor", func(hex string) { win.SetTextColor(hex) })
	w.Bind("gonomoSetTitlebarVisible", func(visible bool) { win.SetTitlebarVisible(visible) })
	w.Bind("gonomoSetTitleBarOverlay", func(enabled bool) { win.SetTitleBarOverlay(enabled) })
	w.Bind("gonomoDismissSplash", func() {
		splash.Dismiss()
		win.Reveal()
	})

	bw, bh := window.CaptionButtonBaseSizes()
	topBorder := win.TopResizeBorderHeight()
	w.Init(fmt.Sprintf(`window.gonomo={
  Config:%s,
  Titlebar:{mode:%q,darkMode:%t,captionColor:"%s",textColor:"%s",titleBarOverlay:%t},
  captionButtonHeight:%s,
  captionButtonsWidth:%s,
  resizeBorderTop:%s,
  Close:window.gonomoClose,
  Minimize:window.gonomoMinimize,
  Maximize:window.gonomoMaximize,
  Restore:window.gonomoRestore,
  IsMaximized:window.gonomoIsMaximized,
  setDarkMode:window.gonomoSetDarkMode,
  setCaptionColor:window.gonomoSetCaptionColor,
  setTextColor:window.gonomoSetTextColor,
  setTitlebarVisible:window.gonomoSetTitlebarVisible,
  setTitleBarOverlay:window.gonomoSetTitleBarOverlay,
  dismissSplash:window.gonomoDismissSplash
};`,
		string(cfgJSON),
		titlebarMode(cfg),
		titlebarDarkMode(cfg),
		titlebarCaptionColor(cfg),
		titlebarTextColor(cfg),
		cfg.UsesTitleBarOverlay(),
		strconv.Itoa(int(bh)),
		strconv.Itoa(int(bw*3)),
		strconv.Itoa(int(topBorder)),
	))
}

func titlebarMode(cfg *config.Config) string {
	if cfg.UsesTitleBarOverlay() {
		return "onlyNativeButtons"
	}
	if cfg.ShouldHideTitlebar() {
		return "hidden"
	}
	return "normal"
}

func titlebarDarkMode(cfg *config.Config) bool {
	return cfg.Window.DarkMode
}

func titlebarCaptionColor(cfg *config.Config) string {
	if cfg.Window.CaptionColor != "" {
		return cfg.Window.CaptionColor
	}
	return "#1e1e2e"
}

func titlebarTextColor(cfg *config.Config) string {
	if cfg.Window.TextColor != "" {
		return cfg.Window.TextColor
	}
	return "#cdd6f4"
}
