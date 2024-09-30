package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func pixelate(img image.Image, pixelSize int) image.Image {
	bounds := img.Bounds()
	pixelated := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += pixelSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += pixelSize {
			var r, g, b, a uint32
			var count uint32

			for py := y; py < y+pixelSize && py < bounds.Max.Y; py++ {
				for px := x; px < x+pixelSize && px < bounds.Max.X; px++ {
					pr, pg, pb, pa := img.At(px, py).RGBA()
					r += pr
					g += pg
					b += pb
					a += pa
					count++
				}
			}

			r /= count
			g /= count
			b /= count
			a /= count

			avgColor := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			for py := y; py < y+pixelSize && py < bounds.Max.Y; py++ {
				for px := x; px < x+pixelSize && px < bounds.Max.X; px++ {
					pixelated.Set(px, py, avgColor)
				}
			}
		}
	}

	saveImage(pixelated, "pixelated.png")

	return pixelated
}

func openImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Decode(file)
	case ".png":
		return png.Decode(file)
	default:
		log.Fatalf("Unsupported file format: %s", ext)
		return nil, nil
	}
}

func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
