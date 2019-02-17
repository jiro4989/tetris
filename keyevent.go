package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

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
				if currentMino.canMoveLeft(currentBoard) {
					currentMino.moveLeft()
				}
			case 'j':
				if currentMino.canMoveDown(currentBoard) {
					currentMino.moveDown()
				}
			case 'l':
				if currentMino.canMoveRight(currentBoard) {
					currentMino.moveRight()
				}
			case 'r':
				// TODO
				fmt.Println("hello")
			case 'e':
				// TODO
				fmt.Println("hello")
			}
		}
		updateDisplayBoard(currentMino)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawBackground()
	}
}
