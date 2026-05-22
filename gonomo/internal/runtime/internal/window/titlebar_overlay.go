package window

import "unsafe"

var (
	titleBarOverlayButtonWidth  = int32(46)
	titleBarOverlayButtonHeight = int32(32)
	titleBarOverlayButtonCount  = int32(3)
	titleBarOverlaySeamOverlap  = int32(1)

	titleBarOverlayNoButton    = -1
	titleBarOverlayMinButton   = 0
	titleBarOverlayMaxButton   = 1
	titleBarOverlayCloseButton = 2
)

type titleBarOverlayState struct {
	hover    int
	pressed  int
	tracking bool
}

var titleBarOverlayParents = map[uintptr]*titleBarOverlayState{}

func SetCaptionButtonSizes(width, height int32) {
	if width >= 46 {
		titleBarOverlayButtonWidth = width
	}
	if height >= 32 {
		titleBarOverlayButtonHeight = height
	}
}

func (w *Window) enableTitleBarOverlayHitTesting() {
	titleBarOverlayParents[w.hwnd] = &titleBarOverlayState{
		hover:   titleBarOverlayNoButton,
		pressed: titleBarOverlayNoButton,
	}

	style, _, _ := procGetWindowLongPtrW.Call(w.hwnd, uintptr(GWL_STYLE))
	style |= uintptr(WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX)
	procSetWindowLongPtrW.Call(w.hwnd, uintptr(GWL_STYLE), style)
	procSetWindowPos.Call(w.hwnd, 0, 0, 0, 0, 0,
		SWP_FRAMECHANGED|SWP_NOMOVE|SWP_NOSIZE)

	applyTitleBarOverlayWidgetRegion(w.hwnd)
	paintTitleBarOverlayButtonsNow(w.hwnd)
}

func (w *Window) disableTitleBarOverlayHitTesting() {
	delete(titleBarOverlayParents, w.hwnd)
	clearTitleBarOverlayWidgetRegion(w.hwnd)
}

func handleTitleBarOverlayHostMessage(hwnd uintptr, msg uint32, wParam uintptr, lParam uintptr) (bool, uintptr) {
	switch msg {
	case WM_NCCALCSIZE:
		if wParam != 0 {
			lParamPtr := *(*unsafe.Pointer)(unsafe.Pointer(&lParam))
			params := (*NCCALCSIZE_PARAMS)(lParamPtr)
			originalTop := params.Rgrc[0].Top

			ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)

			zoomed, _, _ := procIsZoomed.Call(hwnd)
			if zoomed != 0 {
				monitor, _, _ := procMonitorFromWindow.Call(hwnd, MONITOR_DEFAULTTONEAREST)
				var info MONITORINFO
				info.Size = uint32(unsafe.Sizeof(info))
				procGetMonitorInfoW.Call(monitor, uintptr(unsafe.Pointer(&info)))
				params.Rgrc[0].Top = info.Work.Top
			} else {
				params.Rgrc[0].Top = originalTop
			}

			return true, ret
		}
	case WM_NCHITTEST:
		return true, titleBarOverlayHitTest(hwnd, lParam)
	case WM_NCACTIVATE:
		paintTitleBarOverlayButtonsNow(hwnd)
		return true, 1
	case WM_ERASEBKGND:
		eraseTitleBarOverlayBackground(hwnd, wParam)
		return true, 1
	case WM_MOUSEMOVE:
		return true, handleTitleBarOverlayMouseMove(hwnd, lParam)
	case WM_MOUSELEAVE:
		return true, handleTitleBarOverlayMouseLeave(hwnd)
	case WM_LBUTTONDOWN:
		return true, handleTitleBarOverlayLButtonDown(hwnd, lParam)
	case WM_LBUTTONUP:
		return true, handleTitleBarOverlayLButtonUp(hwnd, lParam)
	case WM_NCMOUSEMOVE:
		if wParam == HTMAXBUTTON {
			setTitleBarOverlayHover(hwnd, titleBarOverlayMaxButton, true)
			trackTitleBarOverlayMouseLeave(hwnd, true)
			ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
			paintTitleBarOverlayButtonsNow(hwnd)
			return true, ret
		}
	case WM_NCMOUSELEAVE:
		return true, handleTitleBarOverlayMouseLeave(hwnd)
	case WM_NCLBUTTONDOWN:
		if wParam == HTMAXBUTTON {
			setTitleBarOverlayPressed(hwnd, titleBarOverlayMaxButton)
			return true, 0
		}
	case WM_NCLBUTTONUP:
		state := titleBarOverlayParents[hwnd]
		if state != nil && state.pressed == titleBarOverlayMaxButton {
			state.pressed = titleBarOverlayNoButton
			if titleBarOverlayButtonAt(hwnd, lParam) == HTMAXBUTTON {
				toggleTitleBarOverlayMaximize(hwnd)
			}
			paintTitleBarOverlayButtonsNow(hwnd)
			return true, 0
		}
	case WM_PAINT:
		drawTitleBarOverlayButtons(hwnd)
		return true, 0
	case WM_PARENTNOTIFY:
		applyTitleBarOverlayWidgetRegion(hwnd)
	}
	return false, 0
}
