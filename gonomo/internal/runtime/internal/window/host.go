package window

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var hostClassRegistered bool

func CreateHostWindow(width, height int) (uintptr, error) {
	inst, _, _ := procGetModuleHandleW.Call(0)

	if !hostClassRegistered {
		wc := WNDCLASSEX{
			Size:      uint32(unsafe.Sizeof(WNDCLASSEX{})),
			WndProc:   windows.NewCallback(hostWndProc),
			Instance:  inst,
			ClassName: windows.StringToUTF16Ptr("WebViewHost"),
		}
		ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
		if ret == 0 {
			return 0, fmt.Errorf("RegisterClassExW failed")
		}
		hostClassRegistered = true
	}

	sw, _, _ := procGetSystemMetrics.Call(0)
	sh, _, _ := procGetSystemMetrics.Call(1)
	x := (int(sw) - width) / 2
	y := (int(sh) - height) / 2
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	hwnd, _, _ := procCreateWindowExW.Call(
		0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("WebViewHost"))),
		0,
		WS_OVERLAPPEDWINDOW,
		uintptr(x), uintptr(y),
		uintptr(width), uintptr(height),
		0, 0, inst, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("CreateWindowExW failed")
	}

	return hwnd, nil
}

func hostWndProc(hwnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if titleBarOverlayParents[hwnd] != nil {
		if handled, result := handleTitleBarOverlayHostMessage(hwnd, msg, wParam, lParam); handled {
			return result
		}
	}

	switch msg {
	case WM_SIZE:
		w := int32(lParam & 0xFFFF)
		h := int32(lParam >> 16)
		widget := findWidget(hwnd)
		if widget != 0 && titleBarOverlayParents[hwnd] != nil {
			layoutTitleBarOverlayWidget(hwnd, widget)
		} else if widget != 0 {
			procMoveWindow.Call(widget, 0, 0, uintptr(w), uintptr(h), 1)
		}
		if titleBarOverlayParents[hwnd] != nil {
			applyTitleBarOverlayWidgetRegion(hwnd)
			procInvalidateRect.Call(hwnd, 0, 1)
			procUpdateWindow.Call(hwnd)
		}
		return 0

	case WM_CLOSE:
		procDestroyWindow.Call(hwnd)
		return 0

	case WM_DESTROY:
		delete(titleBarOverlayParents, hwnd)
		clearTitleBarOverlayWidgetRegion(hwnd)
		procPostQuitMessage.Call(0)
		return 0
	}
	ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret
}

func findWidget(hwnd uintptr) uintptr {
	w, _, _ := procFindWindowExW.Call(hwnd, 0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("webview_widget"))), 0)
	return w
}
