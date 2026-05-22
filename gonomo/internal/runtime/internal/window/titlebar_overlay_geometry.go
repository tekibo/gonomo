package window

import "unsafe"

func titleBarOverlayHitTest(hwnd uintptr, lParam uintptr) uintptr {
	switch titleBarOverlayButtonIndexAtScreen(hwnd, lParam) {
	case titleBarOverlayMaxButton:
		return HTMAXBUTTON
	case titleBarOverlayMinButton, titleBarOverlayCloseButton:
		return HTCLIENT
	}
	if hit := titleBarOverlayResizeHitTest(hwnd, lParam); hit != HTCLIENT {
		return hit
	}
	return HTCLIENT
}

func titleBarOverlayButtonAt(hwnd uintptr, lParam uintptr) uintptr {
	switch titleBarOverlayButtonIndexAtScreen(hwnd, lParam) {
	case titleBarOverlayMinButton:
		return HTMINBUTTON
	case titleBarOverlayMaxButton:
		return HTMAXBUTTON
	case titleBarOverlayCloseButton:
		return HTCLOSE
	default:
		return 0
	}
}

func titleBarOverlayButtonIndexAtScreen(hwnd uintptr, lParam uintptr) int {
	x, y := screenPointFromLParam(lParam)
	left, top, right, bottom, ok := titleBarOverlayButtonsRect(hwnd)
	if !ok || x < left || x >= right || y < top || y >= bottom {
		return titleBarOverlayNoButton
	}
	buttonWidth := scaleTitleBarOverlayValue(hwnd, titleBarOverlayButtonWidth)
	switch (x - left) / buttonWidth {
	case 0:
		return titleBarOverlayMinButton
	case 1:
		return titleBarOverlayMaxButton
	case 2:
		return titleBarOverlayCloseButton
	default:
		return titleBarOverlayNoButton
	}
}

func titleBarOverlayButtonIndexAtClient(hwnd uintptr, lParam uintptr) int {
	x, y := pointFromLParam(lParam)
	buttons, ok := titleBarOverlayButtonsClientRect(hwnd)
	if !ok || x < buttons.Left || x >= buttons.Right || y < buttons.Top || y >= buttons.Bottom {
		return titleBarOverlayNoButton
	}
	buttonWidth := (buttons.Right - buttons.Left) / titleBarOverlayButtonCount
	if buttonWidth <= 0 {
		return titleBarOverlayNoButton
	}
	button := int((x - buttons.Left) / buttonWidth)
	if button < titleBarOverlayMinButton || button > titleBarOverlayCloseButton {
		return titleBarOverlayNoButton
	}
	return button
}

func titleBarOverlayButtonsRect(hwnd uintptr) (int32, int32, int32, int32, bool) {
	visible, ok := titleBarOverlayVisibleClientRect(hwnd)
	if !ok {
		return 0, 0, 0, 0, false
	}
	pt := POINT{X: 0, Y: 0}
	if ok, _, _ := procClientToScreen.Call(hwnd, uintptr(unsafe.Pointer(&pt))); ok == 0 {
		return 0, 0, 0, 0, false
	}
	width := scaleTitleBarOverlayValue(hwnd, titleBarOverlayButtonWidth) * titleBarOverlayButtonCount
	height := scaleTitleBarOverlayValue(hwnd, titleBarOverlayButtonHeight)
	left := pt.X + visible.Right - width
	right := pt.X + visible.Right
	top := pt.Y + visible.Top
	bottom := top + height
	return left, top, right, bottom, true
}

func titleBarOverlayButtonsClientRect(hwnd uintptr) (RECT, bool) {
	visible, ok := titleBarOverlayVisibleClientRect(hwnd)
	if !ok {
		return RECT{}, false
	}
	width := scaleTitleBarOverlayValue(hwnd, titleBarOverlayButtonWidth) * titleBarOverlayButtonCount
	height := scaleTitleBarOverlayValue(hwnd, titleBarOverlayButtonHeight)
	left := visible.Right - width
	if left < 0 {
		left = 0
	}
	if visible.Top+height > visible.Bottom {
		height = visible.Bottom - visible.Top
	}
	return RECT{Left: left, Top: visible.Top, Right: visible.Right, Bottom: visible.Top + height}, true
}

func titleBarOverlayVisibleClientRect(hwnd uintptr) (RECT, bool) {
	var client RECT
	if ok, _, _ := procGetClientRect.Call(hwnd, uintptr(unsafe.Pointer(&client))); ok == 0 {
		return RECT{}, false
	}
	return client, true
}

func titleBarOverlayResizeHitTest(hwnd uintptr, lParam uintptr) uintptr {
	zoomed, _, _ := procIsZoomed.Call(hwnd)
	if zoomed != 0 {
		return HTCLIENT
	}

	var rect RECT
	if ok, _, _ := procGetWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&rect))); ok == 0 {
		return HTCLIENT
	}

	x, y := screenPointFromLParam(lParam)
	frameX, _, _ := procGetSystemMetrics.Call(SM_CXFRAME)
	frameY, _, _ := procGetSystemMetrics.Call(SM_CYFRAME)
	padded, _, _ := procGetSystemMetrics.Call(SM_CXPADDEDBORDER)
	borderX := int32(frameX + padded)
	borderY := int32(frameY + padded)

	onLeft := x >= rect.Left && x < rect.Left+borderX
	onRight := x < rect.Right && x >= rect.Right-borderX
	onTop := y >= rect.Top && y < rect.Top+borderY
	onBottom := y < rect.Bottom && y >= rect.Bottom-borderY

	switch {
	case onTop && onLeft:
		return HTTOPLEFT
	case onTop && onRight:
		return HTTOPRIGHT
	case onBottom && onLeft:
		return HTBOTTOMLEFT
	case onBottom && onRight:
		return HTBOTTOMRIGHT
	case onLeft:
		return HTLEFT
	case onRight:
		return HTRIGHT
	case onTop:
		return HTTOP
	case onBottom:
		return HTBOTTOM
	default:
		return HTCLIENT
	}
}

func titleBarOverlayResizeBorder(hwnd uintptr) (int32, int32) {
	zoomed, _, _ := procIsZoomed.Call(hwnd)
	if zoomed != 0 {
		return 0, 0
	}
	frameX, _, _ := procGetSystemMetrics.Call(SM_CXFRAME)
	frameY, _, _ := procGetSystemMetrics.Call(SM_CYFRAME)
	padded, _, _ := procGetSystemMetrics.Call(SM_CXPADDEDBORDER)
	return int32(frameX + padded), int32(frameY + padded)
}

func scaleTitleBarOverlayValue(hwnd uintptr, value int32) int32 {
	dpi, _, _ := procGetDpiForWindow.Call(hwnd)
	if dpi == 0 {
		return value
	}
	return int32(int64(value) * int64(dpi) / USER_DEFAULT_SCREEN_DPI)
}

func isWindowMaximized(hwnd uintptr) bool {
	zoomed, _, _ := procIsZoomed.Call(hwnd)
	return zoomed != 0
}

func screenPointFromLParam(lParam uintptr) (int32, int32) {
	return pointFromLParam(lParam)
}

func pointFromLParam(lParam uintptr) (int32, int32) {
	return int32(int16(lParam & 0xffff)), int32(int16((lParam >> 16) & 0xffff))
}

func maxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
