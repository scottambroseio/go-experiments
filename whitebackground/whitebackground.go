package whitebackground

import (
	"image"
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	go isTopEdgeWhite(ctx, i, c1, x1, x2, y1)
	go isBottomEdgeWhite(ctx, i, c2, x1, x2, y2)
	go isLeftEdgeWhite(ctx, i, c3, x1, y1, y2)
	go isRightEdgeWhite(ctx, i, c4, x2, y1, y2)

	// The moment any channel finds a non white pixel and returns false
	// Cancel the other go routines as there's no longer any need
	// for them to continue
	for count := 0; count < 4; count++ {
		select {
			 case res := <-c1:
				 if !res {
					cancel()
					return false
				 }
			 case res := <-c2:
				 if !res {
					cancel()
					return false
				 }
			 case res := <-c3:
				 if !res {
					cancel()
					return false
				 }
			 case res := <-c4:
				 if !res {
					cancel()
					return false
				 }
		}
	}

	return true
}

func isTopEdgeWhite(ctx context.Context, i image.Image, c chan bool, x1, x2, y int) {
	select {
		case <-ctx.Done(): return
	default:
		for x1 < x2 {
			if !isWhitePixel(i.At(x1, y).RGBA()) {
				c <- false
				return
			}
			x1++
		}
		c <- true
	}
}

func isBottomEdgeWhite(ctx context.Context, i image.Image, c chan bool, x1, x2, y int) {
	select {
		case <-ctx.Done(): return
	default:
		for x1 < x2 {
			if !isWhitePixel(i.At(x1, y).RGBA()) {
				c <- false
				return
			}
			x1++
		}
		c <- true
	}
}

func isLeftEdgeWhite(ctx context.Context, i image.Image, c chan bool, x, y1, y2 int) {
	select {
		case <-ctx.Done(): return
	default:
		for y1 < y2 {
			if !isWhitePixel(i.At(x, y1).RGBA()) {
				c <- false
				return
			}
			y1++
		}
		c <- true
	}
}

func isRightEdgeWhite(ctx context.Context, i image.Image, c chan bool, x, y1, y2 int) {
	select {
		case <-ctx.Done(): return
	default:
		for y1 < y2 {
			if !isWhitePixel(i.At(x, y1).RGBA()) {
				c <- false
				return
			}
			y1++
		}
		c <- true
	}
}

func isWhitePixel(r, g, b, a uint32) bool {
	return r&g&b&a == 65535
}
