package main

//110 x 40 Terminal size
import (
	GoSnake "GoSnake/src"
	"fmt"
	"sync"

	"github.com/nsf/termbox-go"
)

func main() {
	// f := GoSnake.Banner
	MenuSelector(0)
	// err := termbox.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termbox.Close()

	// keyPress := make(chan termbox.Event)

	// go func() {
	// 	for {
	// 		keySeq := termbox.PollEvent()
	// 		keyPress <- keySeq
	// 	}
	// }()

	// //func generate menu
	// // go GetMenu(){

	// // }()

	// for {
	// 	keySeq := <-keyPress
	// 	if keySeq.Type == termbox.EventKey {
	// 		if keySeq.Ch == 'h' {
	// 			fmt.Println("Hello World!")
	// 			// DisplayMenu()

	// 		} else if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
	// 			break
	// 		}
	// 	}
	// }

}

type Options struct {
	display [][]string
	active  int
}

func MenuSelector(preselect int) {

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

	playOption := []string{
		"   ___  __         ",
		"  / _ \\/ /__ ___ __",
		" / ___/ / _ `/ // /",
		"/_/  /_/\\_,_/\\_, / ",
		"            /___/  ",
	}
	levelOption := []string{
		"   __                __",
		"  / /  ___ _  _____ / /",
		" / /__/ -_) |/ / -_) / ",
		"/____/\\__/|___/\\__/_/  ",
	}

	highscoresOption := []string{
		"   __ ___      __     ____                   ",
		"  / // (_)__ _/ /    / __/______  _______ ___",
		" / _  / / _ `/ _ \\  _\\ \\/ __/ _ \\/ __/ -_|_-<",
		"/_//_/_/\\_, /_//_/ /___/\\__/\\___/_/  \\__/___/",
		"       /___/                                 ",
	}

	quitOption := []string{
		"  ____       _ __ ",
		" / __ \\__ __(_) /_",
		"/ /_/ / // / / __/",
		"\\___\\_\\_,_/_/\\__/ ",
	}

	var menuOptions Options
	menuOptions.display = append(menuOptions.display, playOption, levelOption, highscoresOption, quitOption)

	menuOptions.active = preselect

	banner := []string{
		"╔══════════════════════════════════════════════════════╗",
		"║       _____          _____             _             ║",
		"║      / ____|        / ____|           | |            ║",
		"║     | |  __  ___   | (___  _ __   __ _| | _____      ║",
		"║     | | |_ |/ _ \\   \\___ \\| '_ \\ / _` | |/ / _ \\     ║",
		"║     | |__| | (_) |  ____) | | | | (_| |   <  __/     ║",
		"║      \\_____|\\___/  |_____/|_| |_|\\__,_|_|\\_\\___|     ║",
		"║                                                      ║",
		"╚══════════════════════════════════════════════════════╝ ",
	}
	terminalWidth, terminalHeight := termbox.Size()

	bannerPosY := (terminalHeight / 8) - 2
	bannerPosX := (terminalWidth / 2) - (len(banner[1]) / 2)
	dynamicPosY := bannerPosY
	for _, line := range banner {
		dynamicPosX := bannerPosX
		for _, char := range line {
			termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorGreen, termbox.ColorDefault)
			dynamicPosX++
		}
		dynamicPosY++
	}

	optionPosY := dynamicPosY + 1
	optionPosX := (terminalWidth / 2)
	var optionsCords [][]int
	for i, option := range menuOptions.display {
		for i2, line := range option {
			optionPosX = (terminalWidth / 2) - (len(line) / 2)
			if i2 == 0 {
				optionsCords = append(optionsCords, []int{optionPosX, optionPosY, len(line)})
			}
			for _, char := range line {
				if i == preselect {
					termbox.SetCell(optionPosX, optionPosY, char, termbox.ColorRed, termbox.ColorDefault)
					optionPosX++
				} else {
					termbox.SetCell(optionPosX, optionPosY, char, termbox.ColorDefault, termbox.ColorDefault)
					optionPosX++
				}

			}
			optionPosY++
		}
		optionPosY++
	}
	optionSelector(preselect, optionsCords)
	termbox.Sync()

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				termbox.Close()
				break
			}

			if keySeq.Ch == 'w' {
				fmt.Println("wtf")
				optionDeselector(optionsCords)
				termbox.Sync()
			} else if keySeq.Ch == 's' {

			}
		}
	}

}

func optionSelector(active int, optionsPos [][]int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// go func() {
	// 	defer wg.Done() // Notify the wait group when this goroutine is done
	// 	selectorL := GoSnake.ActiveOption[0]

	// 	for i := 0; i < 4; i++ {

	// 		initialX := optionsPos[i][0] - 5
	// 		initialY := optionsPos[i][1]

	// 		for _, line := range selectorL {
	// 			activeX := initialX
	// 			for _, ch := range line {
	// 				ch = ' '
	// 				termbox.SetCell(activeX, initialY, ch, termbox.ColorDefault, termbox.ColorDefault)
	// 				activeX++
	// 			}
	// 			initialY++
	// 		}
	// 	}

	// }()

	go func() {
		defer wg.Done()
		selectorR := GoSnake.ActiveOption[1]
		initialX := optionsPos[active][0] + optionsPos[active][2] + 1
		initialY := optionsPos[active][1]

		for _, line := range selectorR {
			activeX := initialX
			for _, ch := range line {
				termbox.SetCell(activeX, initialY, ch, termbox.ColorDefault, termbox.ColorDefault)
				activeX++
			}
			initialY++
		}
	}()

	go func() {
		defer wg.Done()
		selectorL := GoSnake.ActiveOption[0]
		initialX := optionsPos[active][0] - 5
		initialY := optionsPos[active][1]

		for _, line := range selectorL {
			activeX := initialX
			for _, ch := range line {
				termbox.SetCell(activeX, initialY, ch, termbox.ColorDefault, termbox.ColorDefault)
				activeX++
			}
			initialY++
		}
	}()

	// termbox.SetCell(optionsPos[1][0], optionsPos[1][1], 'X', termbox.ColorDefault, termbox.ColorDefault)
	// termbox.SetCell(optionsPos[2][0], optionsPos[2][1], 'X', termbox.ColorDefault, termbox.ColorDefault)
	// termbox.SetCell(optionsPos[3][0], optionsPos[3][1], 'X', termbox.ColorDefault, termbox.ColorDefault)
	wg.Wait()
}

func optionDeselector(optionsPos [][]int) {
	for i := 0; i >= len(optionsPos); i++ {
		initialX := optionsPos[i][0] - 5
		initialY := optionsPos[i][1]

		for j := 0; j >= 3; j++ {
			for k := 0; k >= 4; k++ {
				termbox.SetCell(initialX+k, initialY+j, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}

}

// selectorR := GoSnake.ActiveOption[1]
// initialX := optionsPos[active][0] + optionsPos[active][2] + 1
// initialY := optionsPos[active][1]

// selectorL := GoSnake.ActiveOption[0]
// initialX := optionsPos[active][0] - 5
// initialY := optionsPos[active][1]
