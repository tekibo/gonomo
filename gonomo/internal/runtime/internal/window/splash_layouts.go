package window

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func drawLayoutCentered(hdc uintptr, s *SplashWindow, cw, ch int32) {
	var graphics uintptr
	if s.imageBitmap != 0 && s.imageWidth > 0 && s.imageHeight > 0 {
		procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
		if graphics != 0 {
			maxH := int32(float64(ch) * 0.55)
			maxW := cw - 40
			dstW, dstH := fitProportional(s.imageWidth, s.imageHeight, maxW, maxH)
			xOff := (cw - dstW) / 2
			yOff := int32(float64(ch) * 0.08)
			procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
				uintptr(xOff), uintptr(yOff), uintptr(dstW), uintptr(dstH))
			procGdipDeleteGraphics.Call(graphics)
		}
	}
	drawTextCentered(hdc, s, cw, ch, 20, int32(float64(ch)*0.60), ch-20)
}

func drawLayoutMinimal(hdc uintptr, s *SplashWindow, cw, ch int32) {
	drawTextCentered(hdc, s, cw, ch, 36, 0, ch)
}

func drawLayoutTopBanner(hdc uintptr, s *SplashWindow, cw, ch int32) {
	if s.imageBitmap != 0 && s.imageWidth > 0 && s.imageHeight > 0 {
		var graphics uintptr
		procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
		if graphics != 0 {
			bannerH := int32(float64(ch) * 0.45)
			dstW, dstH := fitProportional(s.imageWidth, s.imageHeight, cw, bannerH)
			if dstH < bannerH {
				dstH = bannerH
				dstW = int32(float64(s.imageWidth) * float64(bannerH) / float64(s.imageHeight))
			}
			xOff := (cw - dstW) / 2
			procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
				uintptr(xOff), 0, uintptr(dstW), uintptr(dstH))
			procGdipDeleteGraphics.Call(graphics)
		}
	}
	textTop := int32(float64(ch) * 0.50)
	drawTextCentered(hdc, s, cw, ch, 20, textTop, ch-20)
}

func drawLayoutBottomBanner(hdc uintptr, s *SplashWindow, cw, ch int32) {
	textBottom := int32(float64(ch) * 0.50)
	drawTextCentered(hdc, s, cw, ch, 20, 0, textBottom)

	if s.imageBitmap != 0 && s.imageWidth > 0 && s.imageHeight > 0 {
		var graphics uintptr
		procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
		if graphics != 0 {
			bannerH := int32(float64(ch) * 0.45)
			dstW, dstH := fitProportional(s.imageWidth, s.imageHeight, cw, bannerH)
			if dstH < bannerH {
				dstH = bannerH
				dstW = int32(float64(s.imageWidth) * float64(bannerH) / float64(s.imageHeight))
			}
			xOff := (cw - dstW) / 2
			yOff := ch - dstH
			procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
				uintptr(xOff), uintptr(yOff), uintptr(dstW), uintptr(dstH))
			procGdipDeleteGraphics.Call(graphics)
		}
	}
}

func drawLayoutSplit(hdc uintptr, s *SplashWindow, cw, ch int32) {
	halfW := cw / 2
	if s.imageBitmap != 0 && s.imageWidth > 0 && s.imageHeight > 0 {
		var graphics uintptr
		procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
		if graphics != 0 {
			maxW := halfW - 20
			maxH := ch - 40
			dstW, dstH := fitProportional(s.imageWidth, s.imageHeight, maxW, maxH)
			xOff := (halfW - dstW) / 2
			yOff := (ch - dstH) / 2
			procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
				uintptr(xOff), uintptr(yOff), uintptr(dstW), uintptr(dstH))
			procGdipDeleteGraphics.Call(graphics)
		}
	}
	if s.config.Text != "" {
		fontSize := 20
		if ch < 200 {
			fontSize = 14
		}
		font := createSplashFont(fontSize)
		if font != 0 {
			oldFont, _, _ := procSelectObject.Call(hdc, font)
			procSetBkMode.Call(hdc, TRANSPARENT)
			procSetTextColor.Call(hdc, uintptr(toBGR(s.fgColor)))
			var textRC RECT
			textRC.Left = halfW + 10
			textRC.Right = cw - 10
			textRC.Top = 0
			textRC.Bottom = ch
			procDrawTextW.Call(hdc,
				uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(s.config.Text))),
				uintptr(len(s.config.Text)),
				uintptr(unsafe.Pointer(&textRC)),
				DT_CENTER|DT_VCENTER|DT_SINGLELINE)
			if s.config.Tagline != "" {
				tagFont := createSplashFont(fontSize - 6)
				if tagFont != 0 {
					procSelectObject.Call(hdc, tagFont)
					var tagRC RECT
					tagRC.Left = halfW + 10
					tagRC.Right = cw - 10
					tagRC.Top = int32(float64(ch) * 0.55)
					tagRC.Bottom = ch
					procDrawTextW.Call(hdc,
						uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(s.config.Tagline))),
						uintptr(len(s.config.Tagline)),
						uintptr(unsafe.Pointer(&tagRC)),
						DT_CENTER|DT_VCENTER|DT_SINGLELINE)
					procDeleteObject.Call(tagFont)
				}
			}
			procSelectObject.Call(hdc, oldFont)
			procDeleteObject.Call(font)
		}
	}
}

func drawLayoutFullImage(hdc uintptr, s *SplashWindow, cw, ch int32) {
	if s.imageBitmap != 0 && s.imageWidth > 0 && s.imageHeight > 0 {
		var graphics uintptr
		procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
		if graphics != 0 {
			imgAspect := float64(s.imageWidth) / float64(s.imageHeight)
			winAspect := float64(cw) / float64(ch)
			var dstW, dstH int32
			var xOff, yOff int32
			if imgAspect > winAspect {
				dstH = ch
				dstW = int32(float64(dstH) * imgAspect)
				xOff = (cw - dstW) / 2
			} else {
				dstW = cw
				dstH = int32(float64(dstW) / imgAspect)
				yOff = (ch - dstH) / 2
			}
			procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
				uintptr(xOff), uintptr(yOff), uintptr(dstW), uintptr(dstH))
			procGdipDeleteGraphics.Call(graphics)
		}
	}
	drawTextCentered(hdc, s, cw, ch, 24, 0, ch)
}

func drawLayoutCustom(hdc uintptr, s *SplashWindow, cw, ch int32) {
	if s.imageBitmap == 0 || s.imageWidth <= 0 || s.imageHeight <= 0 {
		return
	}
	var graphics uintptr
	procGdipCreateFromHDC.Call(hdc, uintptr(unsafe.Pointer(&graphics)))
	if graphics == 0 {
		return
	}
	dstW, dstH := fitProportional(s.imageWidth, s.imageHeight, cw, ch)
	xOff := (cw - dstW) / 2
	yOff := (ch - dstH) / 2
	procGdipDrawImageRectI.Call(graphics, s.imageBitmap,
		uintptr(xOff), uintptr(yOff), uintptr(dstW), uintptr(dstH))
	procGdipDeleteGraphics.Call(graphics)
}

func fitProportional(imgW, imgH, maxW, maxH int32) (int32, int32) {
	dstW, dstH := imgW, imgH
	if dstW > maxW {
		r := float64(maxW) / float64(dstW)
		dstW = maxW
		dstH = int32(float64(dstH) * r)
	}
	if dstH > maxH {
		r := float64(maxH) / float64(dstH)
		dstH = maxH
		dstW = int32(float64(dstW) * r)
	}
	return dstW, dstH
}

func drawTextCentered(hdc uintptr, s *SplashWindow, cw, ch int32, fontSize int, top, bottom int32) {
	if s.config.Text == "" {
		return
	}
	if ch < 200 && fontSize > 14 {
		fontSize = 14
	}
	font := createSplashFont(fontSize)
	if font == 0 {
		return
	}
	oldFont, _, _ := procSelectObject.Call(hdc, font)
	procSetBkMode.Call(hdc, TRANSPARENT)
	procSetTextColor.Call(hdc, uintptr(toBGR(s.fgColor)))

	var textRC RECT
	textRC.Left = 30
	textRC.Right = cw - 30
	textRC.Top = top
	textRC.Bottom = bottom

	procDrawTextW.Call(hdc,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(s.config.Text))),
		uintptr(len(s.config.Text)),
		uintptr(unsafe.Pointer(&textRC)),
		DT_CENTER|DT_VCENTER|DT_SINGLELINE)

	if s.config.Tagline != "" {
		tagFont := createSplashFont(fontSize - 6)
		if tagFont != 0 {
			procSelectObject.Call(hdc, tagFont)
			var tagRC RECT
			tagRC.Left = 30
			tagRC.Right = cw - 30
			tagRC.Top = top + int32(float64(bottom-top)*0.55)
			tagRC.Bottom = bottom
			procDrawTextW.Call(hdc,
				uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(s.config.Tagline))),
				uintptr(len(s.config.Tagline)),
				uintptr(unsafe.Pointer(&tagRC)),
				DT_CENTER|DT_VCENTER|DT_SINGLELINE)
			procDeleteObject.Call(tagFont)
		}
	}

	procSelectObject.Call(hdc, oldFont)
	procDeleteObject.Call(font)
}

func createSplashFont(size int) uintptr {
	font, _, _ := procCreateFontW.Call(
		uintptr(int32(-size)),
		0, 0, 0,
		FW_BOLD,
		0, 0, 0,
		DEFAULT_CHARSET,
		OUT_DEFAULT_PRECIS,
		CLIP_DEFAULT_PRECIS,
		DEFAULT_QUALITY,
		DEFAULT_PITCH|FF_SWISS,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Segoe UI"))))
	return font
}
