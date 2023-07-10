package GoSnake

import (
	"github.com/nsf/termbox-go"
)

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
	// termbox.SetCell(terminalX-2, terminalY-2, 'X', termbox.ColorDefault, termbox.ColorDefault)
	// termbox.Sync()
	var player snake
	player.headPos = []int{10, 10}
	// updateHeadPos(player, player.headPos[0], player.headPos[1])
	//
	//
	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				termbox.Close()
				break
			}

			if keySeq.Ch == 'w' {
				player.headlastPos = player.headPos
				player.headPos = []int{player.headPos[0], player.headPos[1] - 1}
				updateHeadPos(player, player.headPos[0], player.headPos[1]-1)
			} else if keySeq.Ch == 's' {
				player.headlastPos = player.headPos
				player.headPos = []int{player.headPos[0], player.headPos[1] + 1}
				updateHeadPos(player, player.headPos[0], player.headPos[1]+1)
			} else if keySeq.Ch == 'a' {
				player.headlastPos = player.headPos
				player.headPos = []int{player.headPos[0] - 1, player.headPos[1]}
				updateHeadPos(player, player.headPos[0]-1, player.headPos[1])
			} else if keySeq.Ch == 'd' {
				player.headlastPos = player.headPos
				player.headPos = []int{player.headPos[0] + 1, player.headPos[1]}
				updateHeadPos(player, player.headPos[0]+1, player.headPos[1])
			} else if keySeq.Key == termbox.KeyEnter {

			}
		}
	}

}

type snake struct {
	headlastPos []int
	headPos     []int //cords
	tail        [][]int
}

func updateHeadPos(player snake, x, y int) {

	// player.headlastPos = player.headPos
	// player.headPos = []int{x, y}

	termbox.SetCell(player.headlastPos[0], player.headlastPos[1], ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(player.headPos[0], player.headPos[1], 'o', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Sync()

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
