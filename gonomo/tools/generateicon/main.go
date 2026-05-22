//go:build ignore

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	srcPath := flag.String("src", filepath.Join("resources", "gonomo.png"), "source .png path")
	outPath := flag.String("out", filepath.Join("resources", "gonomo.ico"), "output .ico path")
	flag.Parse()

	if err := writeIcon(*srcPath, *outPath); err != nil {
		fmt.Fprintf(os.Stderr, "generate icon: %v\n", err)
		os.Exit(1)
	}
}

func writeIcon(srcPath, outPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	src, err := png.Decode(srcFile)
	if err != nil {
		return err
	}

	icon := resizeNearest(src, 256, 256)

	var pngData bytes.Buffer
	if err := png.Encode(&pngData, icon); err != nil {
		return err
	}

	header := make([]byte, 6)
	binary.LittleEndian.PutUint16(header[0:], 0) // reserved
	binary.LittleEndian.PutUint16(header[2:], 1) // type = ICO
	binary.LittleEndian.PutUint16(header[4:], 1) // image count

	dirEntry := make([]byte, 16)
	dirEntry[0] = 0                                                      // width: 0 means 256
	dirEntry[1] = 0                                                      // height: 0 means 256
	dirEntry[2] = 0                                                      // color count
	dirEntry[3] = 0                                                      // reserved
	binary.LittleEndian.PutUint16(dirEntry[4:], 1)                       // color planes
	binary.LittleEndian.PutUint16(dirEntry[6:], 32)                      // bits per pixel
	binary.LittleEndian.PutUint32(dirEntry[8:], uint32(pngData.Len()))   // image size
	binary.LittleEndian.PutUint32(dirEntry[12:], uint32(len(header)+16)) // image offset

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, chunk := range [][]byte{header, dirEntry, pngData.Bytes()} {
		if _, err := f.Write(chunk); err != nil {
			return err
		}
	}

	return nil
}

func resizeNearest(src image.Image, width, height int) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := src.Bounds()
	srcW := bounds.Dx()
	srcH := bounds.Dy()

	if srcW == width && srcH == height {
		draw.Draw(dst, dst.Bounds(), src, bounds.Min, draw.Src)
		return dst
	}

	for y := 0; y < height; y++ {
		sy := bounds.Min.Y + y*srcH/height
		for x := 0; x < width; x++ {
			sx := bounds.Min.X + x*srcW/width
			dst.Set(x, y, src.At(sx, sy))
		}
	}

	return dst
}
