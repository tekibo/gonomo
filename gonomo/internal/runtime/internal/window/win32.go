package window

import "golang.org/x/sys/windows"

const (
	DWMWA_USE_IMMERSIVE_DARK_MODE  = 20
	DWMWA_CAPTION_COLOR            = 35
	DWMWA_TEXT_COLOR               = 36
	DWMWA_WINDOW_CORNER_PREFERENCE = 33
	DWMWCP_ROUND                   = 2

	GWL_STYLE           = ^uintptr(15)
	GWL_EXSTYLE         = ^uintptr(19)
	WS_CAPTION          = 0x00C00000
	WS_SYSMENU          = 0x00080000
	WS_THICKFRAME       = 0x00040000
	WS_MINIMIZEBOX      = 0x00020000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_OVERLAPPEDWINDOW = 0x00CF0000
	WS_POPUP            = 0x80000000

	WS_EX_TOOLWINDOW = 0x00000080
	WS_EX_TOPMOST    = 0x00000008
	WS_EX_LAYERED    = 0x00080000

	SW_MINIMIZE      = 6
	SW_MAXIMIZE      = 3
	SW_RESTORE       = 9
	SW_SHOW          = 5
	SW_SHOWMAXIMIZED = 3

	WM_CLOSE         = 0x0010
	WM_SIZE          = 0x0005
	WM_DESTROY       = 0x0002
	WM_NCCALCSIZE    = 0x0083
	WM_NCHITTEST     = 0x0084
	WM_NCPAINT       = 0x0085
	WM_NCACTIVATE    = 0x0086
	WM_NCMOUSEMOVE   = 0x00A0
	WM_NCLBUTTONDOWN = 0x00A1
	WM_NCLBUTTONUP   = 0x00A2
	WM_NCMOUSELEAVE  = 0x02A2
	WM_MOUSEMOVE     = 0x0200
	WM_LBUTTONDOWN   = 0x0201
	WM_LBUTTONUP     = 0x0202
	WM_MOUSELEAVE    = 0x02A3
	WM_SETICON       = 0x0080
	WM_DPICHANGED    = 0x02E0
	WM_PAINT         = 0x000F
	WM_ERASEBKGND    = 0x0014
	WM_PARENTNOTIFY  = 0x0210

	HTCLIENT      = 1
	HTCAPTION     = 2
	HTMINBUTTON   = 8
	HTMAXBUTTON   = 9
	HTLEFT        = 10
	HTRIGHT       = 11
	HTTOP         = 12
	HTTOPLEFT     = 13
	HTTOPRIGHT    = 14
	HTBOTTOM      = 15
	HTBOTTOMLEFT  = 16
	HTBOTTOMRIGHT = 17
	HTCLOSE       = 20

	ICON_SMALL      = 0
	ICON_BIG        = 1
	IMAGE_ICON      = 1
	LR_LOADFROMFILE = 0x0010
	LR_DEFAULTSIZE  = 0x0040

	SWP_NOSIZE       = 0x0001
	SWP_NOMOVE       = 0x0002
	SWP_FRAMECHANGED = 0x0020

	RDW_INVALIDATE = 0x0001
	RDW_UPDATENOW  = 0x0100

	LWA_COLORKEY = 0x00000001

	TRANSPARENT = 1

	SM_CXFRAME        = 32
	SM_CYFRAME        = 33
	SM_CXPADDEDBORDER = 92

	MONITOR_DEFAULTTONEAREST = 2
	USER_DEFAULT_SCREEN_DPI  = 96

	RGN_DIFF      = 4
	PS_SOLID      = 0
	TME_LEAVE     = 0x00000002
	TME_NONCLIENT = 0x00000010

	DT_CENTER     = 0x00000001
	DT_VCENTER    = 0x00000004
	DT_SINGLELINE = 0x00000020
	DT_WORDBREAK  = 0x00000010

	DEFAULT_CHARSET     = 1
	FW_BOLD             = 700
	FF_SWISS            = 0x4000
	OUT_DEFAULT_PRECIS  = 0
	CLIP_DEFAULT_PRECIS = 0
	DEFAULT_QUALITY     = 0
	DEFAULT_PITCH       = 0
)

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}

type TRACKMOUSEEVENT struct {
	Size      uint32
	Flags     uint32
	Track     uintptr
	HoverTime uint32
}

type MONITORINFO struct {
	Size    uint32
	Monitor RECT
	Work    RECT
	Flags   uint32
}

type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}

type PAINTSTRUCT struct {
	Hdc        uintptr
	FErase     uint32
	RcPaint    RECT
	FRestore   uint32
	FIncUpdate uint32
	Reserved   [32]byte
}

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread uint32
	SuppressExternalCodecs   uint32
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type NCCALCSIZE_PARAMS struct {
	Rgrc  [3]RECT
	Lppos uintptr
}

var (
	user32   = windows.NewLazySystemDLL("user32.dll")
	dwmapi   = windows.NewLazySystemDLL("dwmapi.dll")
	gdi32    = windows.NewLazySystemDLL("gdi32.dll")
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")
	gdiplus  = windows.NewLazySystemDLL("gdiplus.dll")

	procShowWindow            = user32.NewProc("ShowWindow")
	procIsZoomed              = user32.NewProc("IsZoomed")
	procSetWindowPos          = user32.NewProc("SetWindowPos")
	procMoveWindow            = user32.NewProc("MoveWindow")
	procGetWindowLongPtrW     = user32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtrW     = user32.NewProc("SetWindowLongPtrW")
	procSendMessageW          = user32.NewProc("SendMessageW")
	procLoadImageW            = user32.NewProc("LoadImageW")
	procGetClientRect         = user32.NewProc("GetClientRect")
	procGetModuleHandleW      = kernel32.NewProc("GetModuleHandleW")
	procDefWindowProcW        = user32.NewProc("DefWindowProcW")
	procDwmSetWindowAttribute = dwmapi.NewProc("DwmSetWindowAttribute")
	procInvalidateRect        = user32.NewProc("InvalidateRect")
	procDeleteObject          = gdi32.NewProc("DeleteObject")
	procCreateSolidBrush      = gdi32.NewProc("CreateSolidBrush")
	procFindWindowExW         = user32.NewProc("FindWindowExW")
	procRegisterClassExW      = user32.NewProc("RegisterClassExW")
	procCreateWindowExW       = user32.NewProc("CreateWindowExW")
	procPostQuitMessage       = user32.NewProc("PostQuitMessage")
	procSetWindowLongW        = user32.NewProc("SetWindowLongW")
	procDestroyWindow         = user32.NewProc("DestroyWindow")
	procGetWindowPlacement    = user32.NewProc("GetWindowPlacement")
	procSetWindowPlacement    = user32.NewProc("SetWindowPlacement")
	procGetWindowRect         = user32.NewProc("GetWindowRect")
	procClientToScreen        = user32.NewProc("ClientToScreen")
	procGetDpiForWindow       = user32.NewProc("GetDpiForWindow")
	procMonitorFromWindow     = user32.NewProc("MonitorFromWindow")
	procGetMonitorInfoW       = user32.NewProc("GetMonitorInfoW")
	procSetWindowRgn          = user32.NewProc("SetWindowRgn")
	procSetCapture            = user32.NewProc("SetCapture")
	procReleaseCapture        = user32.NewProc("ReleaseCapture")
	procTrackMouseEvent       = user32.NewProc("TrackMouseEvent")

	procBeginPaint                 = user32.NewProc("BeginPaint")
	procEndPaint                   = user32.NewProc("EndPaint")
	procFillRect                   = user32.NewProc("FillRect")
	procCreateRectRgn              = gdi32.NewProc("CreateRectRgn")
	procCombineRgn                 = gdi32.NewProc("CombineRgn")
	procCreatePen                  = gdi32.NewProc("CreatePen")
	procMoveToEx                   = gdi32.NewProc("MoveToEx")
	procLineTo                     = gdi32.NewProc("LineTo")
	procDrawTextW                  = user32.NewProc("DrawTextW")
	procSetBkMode                  = gdi32.NewProc("SetBkMode")
	procSetTextColor               = gdi32.NewProc("SetTextColor")
	procCreateFontW                = gdi32.NewProc("CreateFontW")
	procSelectObject               = gdi32.NewProc("SelectObject")
	procGetSystemMetrics           = user32.NewProc("GetSystemMetrics")
	procUpdateWindow               = user32.NewProc("UpdateWindow")
	procRedrawWindow               = user32.NewProc("RedrawWindow")
	procLoadCursorW                = user32.NewProc("LoadCursorW")
	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")

	procGdiplusStartup           = gdiplus.NewProc("GdiplusStartup")
	procGdiplusShutdown          = gdiplus.NewProc("GdiplusShutdown")
	procGdipCreateBitmapFromFile = gdiplus.NewProc("GdipCreateBitmapFromFile")
	procGdipDrawImageRectI       = gdiplus.NewProc("GdipDrawImageRectI")
	procGdipCreateFromHDC        = gdiplus.NewProc("GdipCreateFromHDC")
	procGdipDeleteGraphics       = gdiplus.NewProc("GdipDeleteGraphics")
	procGdipDisposeImage         = gdiplus.NewProc("GdipDisposeImage")
	procGdipGetImageWidth        = gdiplus.NewProc("GdipGetImageWidth")
	procGdipGetImageHeight       = gdiplus.NewProc("GdipGetImageHeight")
)
