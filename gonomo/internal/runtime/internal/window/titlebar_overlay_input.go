package window

import "unsafe"

func handleTitleBarOverlayMouseMove(hwnd uintptr, lParam uintptr) uintptr {
	button := titleBarOverlayButtonIndexAtClient(hwnd, lParam)
	// The maximize button is non-client owned so Windows can show Snap Layouts.
	if button == titleBarOverlayMaxButton {
		button = titleBarOverlayNoButton
	}
	setTitleBarOverlayHover(hwnd, button, false)
	trackTitleBarOverlayMouseLeave(hwnd, false)
	return 0
}

func handleTitleBarOverlayMouseLeave(hwnd uintptr) uintptr {
	state := titleBarOverlayParents[hwnd]
	if state == nil {
		return 0
	}
	state.hover = titleBarOverlayNoButton
	state.tracking = false
	paintTitleBarOverlayButtonsNow(hwnd)
	return 0
}

func handleTitleBarOverlayLButtonDown(hwnd uintptr, lParam uintptr) uintptr {
	button := titleBarOverlayButtonIndexAtClient(hwnd, lParam)
	if button != titleBarOverlayMinButton && button != titleBarOverlayCloseButton {
		return 0
	}
	setTitleBarOverlayPressed(hwnd, button)
	procSetCapture.Call(hwnd)
	return 0
}

func handleTitleBarOverlayLButtonUp(hwnd uintptr, lParam uintptr) uintptr {
	state := titleBarOverlayParents[hwnd]
	if state == nil {
		return 0
	}
	pressed := state.pressed
	state.pressed = titleBarOverlayNoButton
	procReleaseCapture.Call()
	if pressed == titleBarOverlayButtonIndexAtClient(hwnd, lParam) {
		switch pressed {
		case titleBarOverlayMinButton:
			procShowWindow.Call(hwnd, SW_MINIMIZE)
		case titleBarOverlayCloseButton:
			procSendMessageW.Call(hwnd, WM_CLOSE, 0, 0)
		}
	}
	paintTitleBarOverlayButtonsNow(hwnd)
	return 0
}

func setTitleBarOverlayHover(hwnd uintptr, hover int, nonClient bool) {
	state := titleBarOverlayParents[hwnd]
	if state == nil || state.hover == hover {
		return
	}
	state.hover = hover
	if !nonClient || hover == titleBarOverlayMaxButton {
		paintTitleBarOverlayButtonsNow(hwnd)
	}
}

func setTitleBarOverlayPressed(hwnd uintptr, pressed int) {
	state := titleBarOverlayParents[hwnd]
	if state == nil {
		return
	}
	state.pressed = pressed
	paintTitleBarOverlayButtonsNow(hwnd)
}

func trackTitleBarOverlayMouseLeave(hwnd uintptr, nonClient bool) {
	state := titleBarOverlayParents[hwnd]
	if state == nil || state.tracking {
		return
	}
	flags := uint32(TME_LEAVE)
	if nonClient {
		flags |= TME_NONCLIENT
	}
	event := TRACKMOUSEEVENT{
		Size:  uint32(unsafe.Sizeof(TRACKMOUSEEVENT{})),
		Flags: flags,
		Track: hwnd,
	}
	procTrackMouseEvent.Call(uintptr(unsafe.Pointer(&event)))
	state.tracking = true
}

func toggleTitleBarOverlayMaximize(hwnd uintptr) {
	if isWindowMaximized(hwnd) {
		procShowWindow.Call(hwnd, SW_RESTORE)
	} else {
		procShowWindow.Call(hwnd, SW_MAXIMIZE)
	}
}
