package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"

	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

var minValue = 125
var maxValue = 225
var rgb = [3]uint8{125, 255, 255}
var currentIndex = 0

func main() {

	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var startTime = time.Now()
var duration = 100 * time.Second

func run(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state
			gtx := app.NewContext(&ops, e)

			// Grab and 'open' the image
			minecraftImage, err := os.Open("./asset/minecraft.jpg")
			if err != nil {
				log.Fatal(err)
			}

			// Decode the image to be able to viewed
			img, err := jpeg.Decode(minecraftImage)
			if err != nil {
				log.Fatal(err)
			}

			// Build and Paint the Image
			imageOp := paint.NewImageOp(img)
			imageOp.Filter = paint.FilterNearest
			imageOp.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			// Because we're in a FrameEvent (loop), FrameEvents are only issued when the Window is resized or the user interacts with the Window
			// So we need to use Animation to allow us to constantly update the frame. This is how we do it
			elapsed := e.Now.Sub(startTime)
			progress := elapsed.Seconds() / duration.Seconds()

			if progress < 1 {
				e.Source.Execute(op.InvalidateCmd{})
			} else {
				progress = 1
			}

			// Build an rect for the color "over" the image
			bounds := image.Rect(img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Max.X, img.Bounds().Max.Y)
			rrect := clip.RRect{Rect: bounds, SE: 5, SW: 5, NW: 5, NE: 5}
			clip.Outline{Path: rrect.Path(&ops)}.Op().Push(&ops)

			// Get the colors, and index. And then fill it in
			nextIndex := (currentIndex + 1) % 3
			rgb[currentIndex] += 1
			rgb[nextIndex] -= 1
			if rgb[currentIndex] == uint8(minValue) || rgb[nextIndex] == uint8(maxValue) {
				currentIndex = int(nextIndex)
			}
			paint.Fill(&ops, color.NRGBA{R: rgb[0], G: rgb[1], B: rgb[2], A: 100})
			paint.PaintOp{}.Add(&ops)

			e.Frame(gtx.Ops)
		}
	}
}
