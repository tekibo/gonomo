package window

import (
	"fmt"
	"unsafe"
)

func (w *Window) SetDarkMode(enabled bool) {
	w.useDarkMode = enabled
	dark := uint32(0)
	if enabled {
		dark = 1
	}
	procDwmSetWindowAttribute.Call(
		w.hwnd, DWMWA_USE_IMMERSIVE_DARK_MODE,
		uintptr(unsafe.Pointer(&dark)), unsafe.Sizeof(dark))
	procInvalidateRect.Call(w.hwnd, 0, 1)
}

func (w *Window) SetCaptionColor(hex string) {
	color := parseHexColor(hex)
	if color == 0 {
		return
	}
	w.captionColor = color
	bgr := toBGR(color)
	procDwmSetWindowAttribute.Call(
		w.hwnd, DWMWA_CAPTION_COLOR,
		uintptr(unsafe.Pointer(&bgr)), unsafe.Sizeof(bgr))
	if titleBarOverlayParents[w.hwnd] != nil {
		procInvalidateRect.Call(w.hwnd, 0, 1)
	}
}

func (w *Window) SetTextColor(hex string) {
	color := parseHexColor(hex)
	if color == 0 {
		return
	}
	w.textColor = color
	bgr := toBGR(color)
	procDwmSetWindowAttribute.Call(
		w.hwnd, DWMWA_TEXT_COLOR,
		uintptr(unsafe.Pointer(&bgr)), unsafe.Sizeof(bgr))
	if titleBarOverlayParents[w.hwnd] != nil {
		procInvalidateRect.Call(w.hwnd, 0, 1)
	}
}

func (w *Window) SetTitlebarVisible(visible bool) {
	w.disableTitleBarOverlayHitTesting()

	style, _, _ := procGetWindowLongPtrW.Call(w.hwnd, uintptr(GWL_STYLE))
	if visible {
		style |= uintptr(WS_CAPTION | WS_SYSMENU)
	} else {
		style &^= uintptr(WS_CAPTION | WS_SYSMENU)
	}
	procSetWindowLongPtrW.Call(w.hwnd, uintptr(GWL_STYLE), style)
	procSetWindowPos.Call(w.hwnd, 0, 0, 0, 0, 0,
		SWP_FRAMECHANGED|SWP_NOMOVE|SWP_NOSIZE)
}

func (w *Window) SetTitleBarOverlay(enabled bool) {
	if enabled {
		w.enableTitleBarOverlayHitTesting()
	} else {
		w.disableTitleBarOverlayHitTesting()
		w.SetTitlebarVisible(true)
	}
}

func applyDWMDefaults(hwnd uintptr) {
	pref := uint32(DWMWCP_ROUND)
	procDwmSetWindowAttribute.Call(
		hwnd, DWMWA_WINDOW_CORNER_PREFERENCE,
		uintptr(unsafe.Pointer(&pref)), unsafe.Sizeof(pref))
	dark := uint32(1)
	procDwmSetWindowAttribute.Call(
		hwnd, DWMWA_USE_IMMERSIVE_DARK_MODE,
		uintptr(unsafe.Pointer(&dark)), unsafe.Sizeof(dark))
}

func toBGR(rgb uint32) uint32 {
	return ((rgb & 0xFF) << 16) | (rgb & 0xFF00) | ((rgb >> 16) & 0xFF)
}

func parseHexColor(hex string) uint32 {
	if len(hex) == 0 {
		return 0
	}
	if hex[0] == '#' {
		hex = hex[1:]
	}
	var r, g, b uint32
	if len(hex) == 6 {
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
		return (r << 16) | (g << 8) | b
	}
	return 0
}
