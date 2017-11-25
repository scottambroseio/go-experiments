package main

import "github.com/nsf/termbox-go"
import "log"
import "time"

func main() {
	err := termbox.Init()
	defer termbox.Close()

	if err != nil {
		log.Fatal(err)
	}

	termbox.SetCell(0, 0, rune('h'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(1, 0, rune('e'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(2, 0, rune('l'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(3, 0, rune('l'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(4, 0, rune('o'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(5, 0, rune(' '), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(6, 0, rune('w'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(7, 0, rune('o'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(8, 0, rune('r'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(9, 0, rune('l'), termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(10, 0, rune('d'), termbox.ColorWhite, termbox.ColorBlack)

	err = termbox.Flush()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
}
