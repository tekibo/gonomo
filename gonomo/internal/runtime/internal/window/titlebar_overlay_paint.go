package window

import "unsafe"

func eraseTitleBarOverlayBackground(hwnd uintptr, hdc uintptr) {
	bg, _ := titleBarOverlayColors(hwnd)
	fillTitleBarOverlayClientArea(hdc, hwnd, bg)
	buttons, ok := titleBarOverlayButtonsClientRect(hwnd)
	if ok {
		fillTitleBarOverlayRect(hdc, titleBarOverlayClipRect(hwnd, buttons), bg)
	}
}

func fillTitleBarOverlayClientArea(hdc uintptr, hwnd uintptr, color uint32) {
	var clientRect RECT
	if ok, _, _ := procGetClientRect.Call(hwnd, uintptr(unsafe.Pointer(&clientRect))); ok != 0 {
		fillTitleBarOverlayRect(hdc, clientRect, color)
	}
}

func drawTitleBarOverlayButtons(hwnd uintptr) {
	var ps PAINTSTRUCT
	hdc, _, _ := procBeginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
	if hdc == 0 {
		return
	}
	defer procEndPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))

	bg, fg := titleBarOverlayColors(hwnd)
	fillTitleBarOverlayClientArea(hdc, hwnd, bg)

	buttons, ok := titleBarOverlayButtonsClientRect(hwnd)
	if !ok {
		return
	}
	fillTitleBarOverlayRect(hdc, titleBarOverlayClipRect(hwnd, buttons), bg)

	buttonWidth := (buttons.Right - buttons.Left) / titleBarOverlayButtonCount
	state := titleBarOverlayParents[hwnd]
	for i := int32(0); i < titleBarOverlayButtonCount; i++ {
		buttonRect := RECT{
			Left:   buttons.Left + i*buttonWidth,
			Top:    buttons.Top,
			Right:  buttons.Left + (i+1)*buttonWidth,
			Bottom: buttons.Bottom,
		}
		buttonBg := bg
		buttonFg := fg
		button := int(i)
		if state != nil {
			if state.pressed == button {
				buttonBg = pressedTitleBarOverlayColor(bg)
			} else if state.hover == button {
				if button == titleBarOverlayCloseButton {
					buttonBg = 0xe81123
					buttonFg = 0xffffff
				} else {
					buttonBg = hoverTitleBarOverlayColor(bg)
				}
			}
		}
		fillTitleBarOverlayRect(hdc, buttonRect, buttonBg)
		drawTitleBarOverlaySymbol(hdc, buttonRect, button, buttonFg, isWindowMaximized(hwnd))
	}
}

func paintTitleBarOverlayButtonsNow(hwnd uintptr) {
	invalidateTitleBarOverlayButtons(hwnd, true)
	procUpdateWindow.Call(hwnd)
}

func invalidateTitleBarOverlayButtons(hwnd uintptr, erase bool) {
	buttons, ok := titleBarOverlayButtonsClientRect(hwnd)
	if !ok {
		return
	}
	eraseFlag := uintptr(0)
	if erase {
		eraseFlag = 1
	}
	rc := titleBarOverlayClipRect(hwnd, buttons)
	procInvalidateRect.Call(hwnd, uintptr(unsafe.Pointer(&rc)), eraseFlag)
}

func titleBarOverlayColors(hwnd uintptr) (uint32, uint32) {
	bg := uint32(0x1e1e2e)
	fg := uint32(0xcdd6f4)
	if current != nil && current.hwnd == hwnd {
		if current.captionColor != 0 {
			bg = current.captionColor
		}
		if current.textColor != 0 {
			fg = current.textColor
		}
	}
	return bg, fg
}

func fillTitleBarOverlayRect(hdc uintptr, rc RECT, color uint32) {
	brush, _, _ := procCreateSolidBrush.Call(uintptr(toBGR(color)))
	if brush == 0 {
		return
	}
	defer procDeleteObject.Call(brush)
	procFillRect.Call(hdc, uintptr(unsafe.Pointer(&rc)), brush)
}

func hoverTitleBarOverlayColor(color uint32) uint32 {
	return mixTitleBarOverlayColor(color, 0xffffff, 18)
}

func pressedTitleBarOverlayColor(color uint32) uint32 {
	return mixTitleBarOverlayColor(color, 0x000000, 22)
}

func mixTitleBarOverlayColor(from uint32, to uint32, percent int) uint32 {
	fr, fg, fb := int((from>>16)&0xff), int((from>>8)&0xff), int(from&0xff)
	tr, tg, tb := int((to>>16)&0xff), int((to>>8)&0xff), int(to&0xff)
	r := fr + ((tr - fr) * percent / 100)
	g := fg + ((tg - fg) * percent / 100)
	b := fb + ((tb - fb) * percent / 100)
	return uint32(r<<16 | g<<8 | b)
}

func drawTitleBarOverlaySymbol(hdc uintptr, rc RECT, button int, color uint32, maximized bool) {
	pen, _, _ := procCreatePen.Call(PS_SOLID, 1, uintptr(toBGR(color)))
	if pen == 0 {
		return
	}
	oldPen, _, _ := procSelectObject.Call(hdc, pen)
	defer func() {
		if oldPen != 0 {
			procSelectObject.Call(hdc, oldPen)
		}
		procDeleteObject.Call(pen)
	}()

	cx := (rc.Left + rc.Right) / 2
	cy := (rc.Top + rc.Bottom) / 2
	scale := scaleTitleBarOverlayValueFromRect(rc, 1)
	s := int32(5) * scale
	switch button {
	case titleBarOverlayMinButton:
		drawTitleBarOverlayLine(hdc, cx-s, cy, cx+s+1, cy)
	case titleBarOverlayMaxButton:
		if maximized {
			drawTitleBarOverlayStrokeRect(hdc, cx-s+1, cy-s+2, cx+s, cy+s+1)
			drawTitleBarOverlayLine(hdc, cx-s+3, cy-s, cx+s+2, cy-s)
			drawTitleBarOverlayLine(hdc, cx+s+2, cy-s, cx+s+2, cy+s-2)
		} else {
			drawTitleBarOverlayStrokeRect(hdc, cx-s, cy-s, cx+s+1, cy+s+1)
		}
	case titleBarOverlayCloseButton:
		drawTitleBarOverlayLine(hdc, cx-s, cy-s, cx+s+1, cy+s+1)
		drawTitleBarOverlayLine(hdc, cx+s, cy-s, cx-s-1, cy+s+1)
	}
}

func scaleTitleBarOverlayValueFromRect(rc RECT, value int32) int32 {
	height := rc.Bottom - rc.Top
	if height <= titleBarOverlayButtonHeight {
		return value
	}
	return value * height / titleBarOverlayButtonHeight
}

func drawTitleBarOverlayStrokeRect(hdc uintptr, left, top, right, bottom int32) {
	drawTitleBarOverlayLine(hdc, left, top, right, top)
	drawTitleBarOverlayLine(hdc, right, top, right, bottom)
	drawTitleBarOverlayLine(hdc, right, bottom, left, bottom)
	drawTitleBarOverlayLine(hdc, left, bottom, left, top)
}

func drawTitleBarOverlayLine(hdc uintptr, x1, y1, x2, y2 int32) {
	procMoveToEx.Call(hdc, uintptr(x1), uintptr(y1), 0)
	procLineTo.Call(hdc, uintptr(x2), uintptr(y2))
}
