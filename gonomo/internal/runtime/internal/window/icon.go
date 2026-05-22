package window

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"

	"gonomo/runtime/internal/resources"
)

func setWindowIcon(hwnd uintptr, iconPath string) {
	resolved := resolveIconPath(iconPath)
	if resolved == "" {
		extracted, err := resources.Extract(resources.IconFile)
		if err != nil {
			return
		}
		resolved = extracted
	}
	hIcon, _, _ := procLoadImageW.Call(0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(resolved))),
		IMAGE_ICON, 0, 0, LR_LOADFROMFILE|LR_DEFAULTSIZE)
	if hIcon == 0 {
		return
	}
	procSendMessageW.Call(hwnd, WM_SETICON, ICON_SMALL, hIcon)
	procSendMessageW.Call(hwnd, WM_SETICON, ICON_BIG, hIcon)
}

func resolveIconPath(iconPath string) string {
	cleaned := filepath.ToSlash(iconPath)
	cleaned = strings.TrimPrefix(cleaned, "/")
	if filepath.IsAbs(cleaned) {
		if _, err := os.Stat(cleaned); err == nil {
			return cleaned
		}
		return ""
	}
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	resolved := filepath.Join(wd, cleaned)
	if _, err := os.Stat(resolved); err == nil {
		return resolved
	}
	ext := strings.ToLower(filepath.Ext(resolved))
	if ext != ".ico" {
		_, _ = fmt.Fprintf(os.Stderr, "warning: icon file %s not found (expected .ico)\n", resolved)
	}
	return ""
}
