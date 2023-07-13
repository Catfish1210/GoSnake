package GoSnake

import (
	"fmt"
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
	actual    [][]int
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

	// Initialize gameboard make it roughly 1/4 or 1/3 of the size of the terminal 110x40
	// display score top left on top of gameboard
	// display top 10 on rightside of the gameboard
	// when game is over and score is in top 10, have a field(bottom of gameboard) to enter name and save it into leaderboard.json
	// and then return to menu
	terminalX, terminalY := termbox.Size()
	DrawGameboard(terminalX, terminalY)

	//Starting pos
	posX, posY := terminalX/2, terminalY/2
	var player snake
	player.head.position = []int{posX, posY}
	player.head.direction = "E"
	player.tail.length = 0
	// player.tail.positions = [][]int{}
	var lastFoodPos []int
	gameSpeed := 150
	go func() {
		// break on gameover (return f(gameover))

		foodPos := GenerateFood(player, lastFoodPos)
		for posX != terminalX {

			termbox.SetCell(posX, posY, '■', termbox.ColorDefault, termbox.ColorDefault)
			//add tail
			for _, tailcords := range player.tail.actual {
				termbox.SetCell(tailcords[0], tailcords[1], '■', termbox.ColorGreen, termbox.ColorDefault)
			}
			//
			termbox.Sync()
			time.Sleep(time.Duration(gameSpeed) * time.Millisecond)
			//remove tail
			for _, tailcords := range player.tail.actual {
				termbox.SetCell(tailcords[0], tailcords[1], ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
			termbox.SetCell(posX, posY, ' ', termbox.ColorDefault, termbox.ColorDefault)
			//

			termbox.Sync()

			if player.head.direction == "N" {
				gameSpeed = 150
				posY--
			}
			if player.head.direction == "W" {
				posX--
				gameSpeed = 60
			}
			if player.head.direction == "E" {
				posX++
				gameSpeed = 60
			}
			if player.head.direction == "S" {
				posY++
				gameSpeed = 150
			}

			player.head.position = []int{posX, posY}
			if player.head.position[0] == foodPos[0] && player.head.position[1] == foodPos[1] {
				foodPos = GenerateFood(player, lastFoodPos)
				lastFoodPos = foodPos
				player.tail.length++
				player.tail.positions = TrimPath(player)
			}
			player.updateTail()

		}
	}()

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
				termbox.Close()
				fmt.Println("Tail Positions:", player.tail.positions)
				fmt.Println("Actual tail: ", player.tail.actual)
				fmt.Println("Food eaten: ", player.tail.length)

				break
			} else if keySeq.Key == termbox.KeyEnter {
			} else if keySeq.Ch == 'r' || keySeq.Ch == 'R' {
				player.tail.length++
			}
		}
	}
}

func (player *snake) updateTail() {
	player.savePath()
	if player.tail.length > 0 {
		length := player.tail.length
		todraw := player.tail.positions[len(player.tail.positions)-length-1:]
		player.tail.actual = todraw[:len(todraw)-1]
	}
}

func (player *snake) savePath() snake {
	player.tail.positions = append(player.tail.positions, player.head.position)
	return *player
}

func TrimPath(player snake) [][]int {
	actualLen := len(player.tail.actual)
	trimmedPath := player.tail.positions[len(player.tail.positions)-actualLen-1:]
	return trimmedPath
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
