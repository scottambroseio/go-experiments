package main

import (
	"log"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var zeroSprite = [5]byte{
	0xF0,
	0x90,
	0x90,
	0x90,
	0xF0,
}

var oneSprite = [5]byte{
	0x20,
	0x60,
	0x20,
	0x20,
	0x70,
}

var twoSprite = [5]byte{
	0xF0,
	0x10,
	0xF0,
	0x80,
	0xF0,
}

var threeSprite = [5]byte{
	0xF0,
	0x10,
	0xF0,
	0x10,
	0xF0,
}

var fourSprite = [5]byte{
	0x90,
	0x90,
	0xF0,
	0x10,
	0x10,
}

var fiveSprite = [5]byte{
	0xF0,
	0x80,
	0xF0,
	0x10,
	0xF0,
}

var sixSprite = [5]byte{
	0xF0,
	0x80,
	0xF0,
	0x90,
	0xF0,
}

var sevenSprite = [5]byte{
	0xF0,
	0x10,
	0x20,
	0x40,
	0x40,
}

var eightSprite = [5]byte{
	0xF0,
	0x90,
	0xF0,
	0x90,
	0xF0,
}

var nineSprite = [5]byte{
	0xF0,
	0x90,
	0xF0,
	0x10,
	0xF0,
}

var sprites = [10][5]byte{zeroSprite, oneSprite, twoSprite, threeSprite, fourSprite, fiveSprite, sixSprite, sevenSprite, eightSprite, nineSprite}

func main() {
	err := termbox.Init()
	defer termbox.Close()

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sprites {
		drawSprite(s)
		err = termbox.Flush()
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func getBitsFromByte(b byte) [8]byte {
	return [8]byte{
		(b << 0) >> 7,
		(b << 1) >> 7,
		(b << 2) >> 7,
		(b << 3) >> 7,
		(b << 4) >> 7,
		(b << 5) >> 7,
		(b << 6) >> 7,
		(b << 7) >> 7,
	}
}

func drawSprite(sprite [5]byte) {
	for i, v := range sprite {
		for i2, v2 := range getBitsFromByte(v) {
			if v2 == 1 {
				termbox.SetCell(i2, i, ' ', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(i2, i, ' ', termbox.ColorWhite, termbox.ColorDefault)
			}
		}
	}
}
