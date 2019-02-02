package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	// termboxの初期化
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	waitKeyInput()
}

func waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD:
				return
			}
			switch ev.Ch {
			case 'q':
				return
			case 'h':
				// TODO
				fmt.Println("hello")
			case 'j':
				// TODO
				fmt.Println("hello")
			case 'l':
				// TODO
				fmt.Println("hello")
			case 'r':
				// TODO
				fmt.Println("hello")
			case 'e':
				// TODO
				fmt.Println("hello")
			}
		}
		// TODO
		// redraw
	}
}
