package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var currentMino Mino

func main() {
	// termboxの初期化
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	currentMino = newMino()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	drawBackground()

	go clock()

	waitKeyInput()
}

func clock() {
	for {
		if !currentMino.canMoveDown(currentBoard) {
			updateCurrentBoard(currentMino)
			currentMino = newMino()
			if !currentMino.canMoveDown(currentBoard) {
				fmt.Println("game over")
			}
		} else {
			currentMino.moveDown()
		}
		drawBackground()
		time.Sleep(1 * time.Second)
	}
}
