package window

import (
	"unsafe"

	webview "github.com/webview/webview_go"
)

type Window struct {
	hwnd           uintptr
	useDarkMode    bool
	captionColor   uint32
	textColor      uint32
	startMaximized bool
}

var current *Window

func Init(w webview.WebView, iconPath string, startMaximized bool) *Window {
	hwnd := uintptr(w.Window())
	win := &Window{hwnd: hwnd, useDarkMode: true, startMaximized: startMaximized}
	current = win
	setWindowIcon(hwnd, iconPath)
	applyDWMDefaults(hwnd)
	return win
}

func Get() *Window { return current }

func (w *Window) Reveal() {
	if w.startMaximized {
		var wp WINDOWPLACEMENT
		wp.Length = uint32(unsafe.Sizeof(wp))
		procGetWindowPlacement.Call(w.hwnd, uintptr(unsafe.Pointer(&wp)))
		wp.ShowCmd = SW_SHOWMAXIMIZED
		procSetWindowPlacement.Call(w.hwnd, uintptr(unsafe.Pointer(&wp)))
	} else {
		procShowWindow.Call(w.hwnd, SW_SHOW)
	}
}

func (w *Window) Minimize() { procShowWindow.Call(w.hwnd, SW_MINIMIZE) }
func (w *Window) Maximize() { procShowWindow.Call(w.hwnd, SW_MAXIMIZE) }
func (w *Window) Restore()  { procShowWindow.Call(w.hwnd, SW_RESTORE) }
func (w *Window) Close()    { procSendMessageW.Call(w.hwnd, WM_CLOSE, 0, 0) }

func (w *Window) IsMaximized() bool {
	ret, _, _ := procIsZoomed.Call(w.hwnd)
	return ret != 0
}

func (w *Window) CaptionButtonSizes() (buttonWidth, buttonHeight, totalWidth int32) {
	bw := scaleTitleBarOverlayValue(w.hwnd, titleBarOverlayButtonWidth)
	bh := scaleTitleBarOverlayValue(w.hwnd, titleBarOverlayButtonHeight)
	return bw, bh, bw * titleBarOverlayButtonCount
}

func CaptionButtonBaseSizes() (buttonWidth, buttonHeight int32) {
	return titleBarOverlayButtonWidth, titleBarOverlayButtonHeight
}

func (w *Window) TopResizeBorderHeight() int32 {
	_, h := titleBarOverlayResizeBorder(w.hwnd)
	return h
}
