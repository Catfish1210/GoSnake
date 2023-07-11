package GoSnake

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type snake struct {
	head head
	tail tail
}

type head struct {
	position  []int
	direction string
}

type tail struct {
	length    int
	positions [][]int
}

func InitializeGame() {
	// Initialize termbox and event listener
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

	// Initialize gameboard
	terminalX, terminalY := termbox.Size()
	DrawGameboard(terminalX, terminalY)

	//Starting pos
	posX, posY := terminalX/2, terminalY/2
	var player snake
	player.head.position = []int{posX, posY}
	player.head.direction = "E"
	// player.tail.length = 3
	// player.tail.positions = [][]int{}
	var lastFoodPos []int
	gameSpeed := 200
	go func() {
		// break on gameover (return f(gameover))

		foodPos := GenerateFood(player, lastFoodPos)
		for posX != terminalX {

			termbox.SetCell(posX, posY, '■', termbox.ColorDefault, termbox.ColorDefault)
			termbox.Sync()
			time.Sleep(time.Duration(gameSpeed) * time.Millisecond)
			termbox.SetCell(posX, posY, ' ', termbox.ColorDefault, termbox.ColorDefault)
			updateTail(player)
			termbox.Sync()

			if player.head.direction == "N" {
				gameSpeed = 310
				posY--
			}
			if player.head.direction == "W" {
				posX--
				gameSpeed = 200
			}
			if player.head.direction == "E" {
				posX++
				gameSpeed = 200
			}
			if player.head.direction == "S" {
				posY++
				gameSpeed = 310
			}

			player.head.position = []int{posX, posY}
			if player.head.position[0] == foodPos[0] && player.head.position[1] == foodPos[1] {
				foodPos = GenerateFood(player, lastFoodPos)
				lastFoodPos = foodPos
			}

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
			} else if keySeq.Ch == 'f' || keySeq.Ch == 'F' {
				foodPos := GenerateFood(player, lastFoodPos)
				lastFoodPos = foodPos
			} else if keySeq.Key == termbox.KeyEnter {

			}
		}
	}

}

func updateTail(player snake) {

	player.tail.positions = append(player.tail.positions, []int{player.head.position[0], player.head.position[1]})
	//
	for _, tails := range player.tail.positions {
		// if player.head.direction == "N" || player.head.direction == "S" {
		// 	termbox.SetCell(tails[0], tails[1], '█', termbox.ColorGreen, termbox.ColorDefault)
		// } else {
		// 	termbox.SetCell(tails[0], tails[1], '■', termbox.ColorGreen, termbox.ColorDefault)
		// }
		termbox.SetCell(tails[0], tails[1], '■', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.Sync()

}

func GenerateFood(player snake, lastFoodPos []int) []int {
	var foodSpawnBlacklist [][]int
	if lastFoodPos != nil {
		foodSpawnBlacklist = append(foodSpawnBlacklist, lastFoodPos)
	}

	terminalX, terminalY := termbox.Size()
	seed := time.Now().UnixNano()
	RandomGenerator := rand.New(rand.NewSource(seed))
	for {
		foodX := RandomGenerator.Intn(terminalX - 1)
		foodY := RandomGenerator.Intn(terminalY - 1)
		if foodX == 0 {
			foodX++
		}
		if foodY == 0 {
			foodY++
		}
		foodPos := []int{foodX, foodY}
		termbox.SetCell(foodX, foodY, '■', termbox.ColorRed, termbox.ColorDefault)
		termbox.Sync()
		return foodPos

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
