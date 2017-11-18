package whitebackground

import (
	"image"
)

// HasWhiteBackground checks to see whether has a white background
// Decided by checking all the edge pixels in turn
func HasWhiteBackground(i image.Image) bool {
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := make(chan bool)
	c4 := make(chan bool)

	b := i.Bounds()

	x1, x2, y1, y2 := b.Min.X, b.Max.X-1, b.Min.Y, b.Max.Y-1

	go isTopEdgeWhite(i, c1, x1, x2, y1)
	go isBottomEdgeWhite(i, c2, x1, x2, y2)
	go isLeftEdgeWhite(i, c3, x1, y1, y2)
	go isRightEdgeWhite(i, c4, x2, y1, y2)

	return <-c1 && <-c2 && <-c3 && <-c4
}

func isTopEdgeWhite(i image.Image, c chan bool, x1, x2, y int) {
	for x1 < x2 {
		if !isWhitePixel(i.At(x1, y).RGBA()) {
			c <- false
			return
		}
		x1++
	}
	c <- true
}

func isBottomEdgeWhite(i image.Image, c chan bool, x1, x2, y int) {
	for x1 < x2 {
		if !isWhitePixel(i.At(x1, y).RGBA()) {
			c <- false
			return
		}
		x1++
	}
	c <- true
}

func isLeftEdgeWhite(i image.Image, c chan bool, x, y1, y2 int) {
	for y1 < y2 {
		if !isWhitePixel(i.At(x, y1).RGBA()) {
			c <- false
			return
		}
		y1++
	}
	c <- true
}

func isRightEdgeWhite(i image.Image, c chan bool, x, y1, y2 int) {
	for y1 < y2 {
		if !isWhitePixel(i.At(x, y1).RGBA()) {
			c <- false
			return
		}
		y1++
	}
	c <- true
}

func isWhitePixel(r, g, b, a uint32) bool {
	return r&g&b&a == 65535
}
