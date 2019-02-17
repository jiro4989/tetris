package main

import (
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

	waitKeyInput()
}
