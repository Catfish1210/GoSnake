package GoSnake

import (
	"time"

	"github.com/nsf/termbox-go"
)

type snake struct {
	head head
}

type head struct {
	position  []int
	direction string
}

func InitializeGame() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	keyPress := make(chan termbox.Event)
	go func() {
		for {
			keySeq := termbox.PollEvent()
			keyPress <- keySeq
		}
	}()
	//
	//
	terminalX, terminalY := termbox.Size()
	DrawGameboard(terminalX, terminalY)

	//Starting pos
	posX, posY := terminalX/2, terminalY/2
	var player snake
	player.head.position = []int{posX, posY}
	player.head.direction = "E"

	go func() {
		for posX != terminalX {
			termbox.SetCell(posX, posY, '■', termbox.ColorDefault, termbox.ColorDefault)
			termbox.Sync()
			time.Sleep(1 * time.Second)
			termbox.SetCell(posX, posY, ' ', termbox.ColorDefault, termbox.ColorDefault)
			termbox.Sync()
			if player.head.direction == "N" {
				posY--
			}
			if player.head.direction == "W" {
				posX--
			}
			if player.head.direction == "E" {
				posX++
			}
			if player.head.direction == "S" {
				posY++
			}

			// posX++
			// posY++

		}
	}()
	//
	//
	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				termbox.Close()
				break
			}

			if keySeq.Ch == 'w' || keySeq.Ch == 'W' {
				player.head.direction = "N"
			} else if keySeq.Ch == 's' || keySeq.Ch == 'S' {
				player.head.direction = "S"
			} else if keySeq.Ch == 'a' || keySeq.Ch == 'A' {
				player.head.direction = "W"
			} else if keySeq.Ch == 'd' || keySeq.Ch == 'D' {
				player.head.direction = "E"
			} else if keySeq.Key == termbox.KeyEnter {

			}
		}
	}

}

func DrawGameboard(terminalX, terminalY int) {
	//getAllBorders (slice of cords) [][]int
	drawElements := BorderPrimitives
	for verticalPos := 0; verticalPos <= terminalY; verticalPos++ {
		for horizontalPos := 0; horizontalPos <= terminalX; horizontalPos++ {

			// Left-Top Corner
			if verticalPos == 0 && horizontalPos == 0 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[2], termbox.ColorGreen, termbox.ColorDefault)
				// Right-Top Corner
			} else if verticalPos == 0 && horizontalPos == terminalX-1 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[3], termbox.ColorGreen, termbox.ColorDefault)
				// Left-Bot Corner
			} else if verticalPos == terminalY-1 && horizontalPos == 0 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[4], termbox.ColorGreen, termbox.ColorDefault)
				// Right-Bot Corner
			} else if verticalPos == terminalY-1 && horizontalPos == terminalX-1 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[5], termbox.ColorGreen, termbox.ColorDefault)
				// Horizontal Border
			} else if verticalPos == 0 || verticalPos == terminalY-1 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[0], termbox.ColorGreen, termbox.ColorDefault)
				// Vertical Border
			} else if horizontalPos == 0 || horizontalPos == terminalX-1 {
				termbox.SetCell(horizontalPos, verticalPos, drawElements[1], termbox.ColorGreen, termbox.ColorDefault)
			}

		}
	}
	termbox.Sync()
}
