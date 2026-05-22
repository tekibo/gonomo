package window

func applyTitleBarOverlayWidgetRegion(hwnd uintptr) {
	widget := findWidget(hwnd)
	if widget == 0 {
		return
	}

	visible, ok := titleBarOverlayVisibleClientRect(hwnd)
	if !ok {
		return
	}
	buttons, ok := titleBarOverlayButtonsClientRect(hwnd)
	if !ok {
		return
	}
	buttonsClip := titleBarOverlayClipRect(hwnd, buttons)
	buttonsClip.Left -= visible.Left
	buttonsClip.Right -= visible.Left
	buttonsClip.Top -= visible.Top
	buttonsClip.Bottom -= visible.Top
	if buttonsClip.Left < 0 {
		buttonsClip.Left = 0
	}
	if buttonsClip.Top < 0 {
		buttonsClip.Top = 0
	}
	widgetWidth := visible.Right - visible.Left
	widgetHeight := visible.Bottom - visible.Top

	fullRgn, _, _ := procCreateRectRgn.Call(0, 0, uintptr(widgetWidth), uintptr(widgetHeight))
	buttonRgn, _, _ := procCreateRectRgn.Call(uintptr(buttonsClip.Left), uintptr(buttonsClip.Top), uintptr(buttonsClip.Right), uintptr(buttonsClip.Bottom))
	if fullRgn == 0 || buttonRgn == 0 {
		if fullRgn != 0 {
			procDeleteObject.Call(fullRgn)
		}
		if buttonRgn != 0 {
			procDeleteObject.Call(buttonRgn)
		}
		return
	}

	procCombineRgn.Call(fullRgn, fullRgn, buttonRgn, RGN_DIFF)
	procDeleteObject.Call(buttonRgn)

	_, borderY := titleBarOverlayResizeBorder(hwnd)
	if borderY > 0 {
		edgeRgn, _, _ := procCreateRectRgn.Call(0, 0, uintptr(widgetWidth), uintptr(borderY))
		if edgeRgn != 0 {
			procCombineRgn.Call(fullRgn, fullRgn, edgeRgn, RGN_DIFF)
			procDeleteObject.Call(edgeRgn)
		}
	}

	// SetWindowRgn owns fullRgn after success.
	if ok, _, _ := procSetWindowRgn.Call(widget, fullRgn, 1); ok == 0 {
		procDeleteObject.Call(fullRgn)
	}
}

func clearTitleBarOverlayWidgetRegion(hwnd uintptr) {
	if widget := findWidget(hwnd); widget != 0 {
		procSetWindowRgn.Call(widget, 0, 1)
	}
}

func layoutTitleBarOverlayWidget(hwnd uintptr, widget uintptr) {
	visible, ok := titleBarOverlayVisibleClientRect(hwnd)
	if !ok {
		return
	}
	procMoveWindow.Call(
		widget,
		uintptr(visible.Left),
		uintptr(visible.Top),
		uintptr(visible.Right-visible.Left),
		uintptr(visible.Bottom-visible.Top),
		1,
	)
}

func titleBarOverlayClipRect(hwnd uintptr, rc RECT) RECT {
	overlap := scaleTitleBarOverlayValue(hwnd, titleBarOverlaySeamOverlap)
	rc.Left -= overlap

	visible, ok := titleBarOverlayVisibleClientRect(hwnd)
	if !ok {
		visible = RECT{}
	}
	if rc.Left < visible.Left {
		rc.Left = visible.Left
	}
	if rc.Top < visible.Top {
		rc.Top = visible.Top
	}
	if visible.Right > visible.Left && rc.Right > visible.Right {
		rc.Right = visible.Right
	}
	if visible.Bottom > visible.Top && rc.Bottom > visible.Bottom {
		rc.Bottom = visible.Bottom
	}
	return rc
}
