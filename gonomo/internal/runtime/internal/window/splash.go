package window

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"

	"gonomo/runtime/internal/config"
	"gonomo/runtime/internal/resources"
)

type SplashWindow struct {
	hwnd         uintptr
	config       config.SplashConfig
	bgBrush      uintptr
	fgColor      uint32
	imageBitmap  uintptr
	imageWidth   int32
	imageHeight  int32
	gpToken      uintptr
	created      time.Time
	resolvedPath string
}

var (
	splashClassRegistered bool
	activeSplash          *SplashWindow
)

func CreateSplash(cfg config.SplashConfig) (*SplashWindow, error) {
	if !cfg.Enabled {
		return nil, nil
	}

	w := cfg.WidthOrDefault()
	h := cfg.HeightOrDefault()

	bgColor := parseHexColor(cfg.BackgroundColor)
	fgColor := parseHexColor(cfg.ForegroundColor)

	inst, _, _ := procGetModuleHandleW.Call(0)

	if !splashClassRegistered {
		className := windows.StringToUTF16Ptr("GoWebViewSplash")
		var wc WNDCLASSEX
		wc.Size = uint32(unsafe.Sizeof(wc))
		wc.WndProc = windows.NewCallback(splashWndProc)
		wc.Instance = inst
		wc.Cursor = loadDefaultCursor()
		wc.ClassName = className
		procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
		splashClassRegistered = true
	}

	sw, _, _ := procGetSystemMetrics.Call(0)
	sh, _, _ := procGetSystemMetrics.Call(1)
	x := (int(sw) - w) / 2
	y := (int(sh) - h) / 2
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	hwnd, _, _ := procCreateWindowExW.Call(
		WS_EX_TOOLWINDOW|WS_EX_TOPMOST,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("GoWebViewSplash"))),
		0,
		WS_POPUP,
		uintptr(x), uintptr(y),
		uintptr(w), uintptr(h),
		0, 0, inst, 0)
	if hwnd == 0 {
		return nil, fmt.Errorf("CreateWindowExW failed")
	}

	bgBrush, _, _ := procCreateSolidBrush.Call(uintptr(toBGR(bgColor)))

	var input GdiplusStartupInput
	input.GdiplusVersion = 1
	var output GdiplusStartupOutput
	var gpToken uintptr
	status, _, _ := procGdiplusStartup.Call(
		uintptr(unsafe.Pointer(&gpToken)),
		uintptr(unsafe.Pointer(&input)),
		uintptr(unsafe.Pointer(&output)))
	if status != 0 {
		procDeleteObject.Call(bgBrush)
		procDestroyWindow.Call(hwnd)
		return nil, fmt.Errorf("GdiplusStartup failed: %d", status)
	}

	var imageBitmap uintptr
	var imgW, imgH int32
	if cfg.Image != "" {
		resolved := resolveSplashImagePath(cfg.Image)
		if resolved != "" {
			u16path := windows.StringToUTF16Ptr(resolved)
			procGdipCreateBitmapFromFile.Call(
				uintptr(unsafe.Pointer(u16path)),
				uintptr(unsafe.Pointer(&imageBitmap)))
			if imageBitmap != 0 {
				var iw, ih int32
				procGdipGetImageWidth.Call(imageBitmap, uintptr(unsafe.Pointer(&iw)))
				procGdipGetImageHeight.Call(imageBitmap, uintptr(unsafe.Pointer(&ih)))
				imgW, imgH = iw, ih
			}
		}
	}

	win := &SplashWindow{
		hwnd:        hwnd,
		config:      cfg,
		bgBrush:     bgBrush,
		fgColor:     fgColor,
		imageBitmap: imageBitmap,
		imageWidth:  imgW,
		imageHeight: imgH,
		gpToken:     gpToken,
		created:     time.Now(),
	}

	if cfg.LayoutOrDefault() == "custom" {
		exStyle, _, _ := procGetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE)
		procSetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE, exStyle|WS_EX_LAYERED)
		procSetLayeredWindowAttributes.Call(hwnd, uintptr(toBGR(bgColor)), 0, LWA_COLORKEY)
	}

	activeSplash = win

	procShowWindow.Call(hwnd, SW_SHOW)
	procUpdateWindow.Call(hwnd)
	procRedrawWindow.Call(hwnd, 0, 0, RDW_INVALIDATE|RDW_UPDATENOW)

	return win, nil
}

func (s *SplashWindow) Dismiss() {
	if s == nil || activeSplash == nil {
		return
	}

	elapsed := time.Since(s.created)
	minDur := time.Duration(s.config.MinDurationOrDefault()) * time.Millisecond
	if elapsed < minDur {
		time.Sleep(minDur - elapsed)
	}

	procDestroyWindow.Call(s.hwnd)

	if s.imageBitmap != 0 {
		procGdipDisposeImage.Call(s.imageBitmap)
	}
	if s.bgBrush != 0 {
		procDeleteObject.Call(s.bgBrush)
	}
	if s.gpToken != 0 {
		procGdiplusShutdown.Call(s.gpToken)
	}

	activeSplash = nil
}

func loadDefaultCursor() uintptr {
	ret, _, _ := procLoadCursorW.Call(0, uintptr(32512))
	return ret
}

func resolveSplashImagePath(imagePath string) string {
	cleaned := filepath.ToSlash(imagePath)
	cleaned = strings.TrimPrefix(cleaned, "/")
	if filepath.IsAbs(cleaned) {
		if _, err := os.Stat(cleaned); err == nil {
			return cleaned
		}
		return ""
	}

	// 1. Try to extract from embedded resources
	if extracted, err := resources.Extract(cleaned); err == nil {
		return extracted
	}

	// 2. Fallback to working directory
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	resolved := filepath.Join(wd, cleaned)
	if _, err := os.Stat(resolved); err == nil {
		return resolved
	}
	return ""
}

func splashWndProc(hwnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case WM_PAINT:
		s := activeSplash
		if s == nil {
			break
		}
		var ps PAINTSTRUCT
		hdc, _, _ := procBeginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
		defer procEndPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))

		var rc RECT
		procGetClientRect.Call(hwnd, uintptr(unsafe.Pointer(&rc)))
		cw := rc.Right - rc.Left
		ch := rc.Bottom - rc.Top

		layout := s.config.LayoutOrDefault()
		if layout != "custom" {
			procFillRect.Call(hdc, uintptr(unsafe.Pointer(&rc)), s.bgBrush)
		}

		switch layout {
		case "minimal":
			drawLayoutMinimal(hdc, s, cw, ch)
		case "top-banner":
			drawLayoutTopBanner(hdc, s, cw, ch)
		case "bottom-banner":
			drawLayoutBottomBanner(hdc, s, cw, ch)
		case "split":
			drawLayoutSplit(hdc, s, cw, ch)
		case "full-image":
			drawLayoutFullImage(hdc, s, cw, ch)
		case "custom":
			drawLayoutCustom(hdc, s, cw, ch)
		default:
			drawLayoutCentered(hdc, s, cw, ch)
		}

		return 0

	case WM_ERASEBKGND:
		return 1
	}

	ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret
}
