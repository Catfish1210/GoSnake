package main

import (
	// GoSnake "GoSnake/src"
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	// f := GoSnake.Banner
	MenuSelector()
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

	//func generate menu
	// go GetMenu(){

	// }()

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Ch == 'h' {
				fmt.Println("Hello World!")
				// DisplayMenu()

			} else if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				break
			}
		}
	}

}

type Options struct {
	display  [][]string
	selector [][]string
	active   int
}

func MenuSelector() {

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

	activeOption := [][]string{
		{
			"__   ",
			"\\ \\  ",
			" > > ",
			"/_/  ",
		},
		{
			"  __ ",
			" / / ",
			"< <  ",
			" \\_\\ ",
		},
	}

	var menuOptions Options
	menuOptions.active = 0
	menuOptions.display = append(menuOptions.display, playOption, levelOption, highscoresOption, quitOption)
	menuOptions.selector = activeOption

	// terminalWidth, terminalHeight := termbox.Size()

	// Display Banner
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

	//apply selector to first option
	// for i, v := range menuOptions.display[menuOptions.active] {
	// 	if menuOptions.active == i {
	// 		combinedLine := menuOptions.selector[0][i] + v + menuOptions.selector[1][i]
	// 		menuOptions.display[menuOptions.active][i] = combinedLine
	// 	}
	// }

	for _, option := range menuOptions.display {
		for _, line := range option {
			optionPosX = (terminalWidth / 2) - (len(line) / 2)
			for _, char := range line {
				termbox.SetCell(optionPosX, optionPosY, char, termbox.ColorDefault, termbox.ColorDefault)
				optionPosX++
			}
			optionPosY++
		}
		optionPosY++
	}

	termbox.Sync()

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Ch == 'h' {
				fmt.Println('h')
			} else if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				break
			}
		}
	}

}

// banner := []string{
// 	"   _____          _____             _         ",
// 	"  / ____|        / ____|           | |        ",
// 	" | |  __  ___   | (___  _ __   __ _| | _____  ",
// 	" | | |_ |/ _ \\   \\___ \\| '_ \\ / _` | |/ / _ \\ ",
// 	" | |__| | (_) |  ____) | | | | (_| |   <  __/ ",
// 	"  \\_____|\\___/  |_____/|_| |_|\\__,_|_|\\_\\___| ",
// }

// func displayBanner() {
// 	banner := []string{
// 		"   _____          _____             _         ",
// 		"  / ____|        / ____|           | |        ",
// 		" | |  __  ___   | (___  _ __   __ _| | _____  ",
// 		" | | |_ |/ _ \\   \\___ \\| '_ \\ / _` | |/ / _ \\ ",
// 		" | |__| | (_) |  ____) | | | | (_| |   <  __/ ",
// 		"  \\_____|\\___/  |_____/|_| |_|\\__,_|_|\\_\\___| ",
// 	}
// 	terminalWidth, terminalHeight := termbox.Size()

// 	// Display Banner
// 	bannerPosY := terminalHeight / 8
// 	bannerPosX := (terminalWidth / 2) - 23
// 	dynamicPosY := bannerPosY
// 	for _, line := range banner {
// 		dynamicPosX := bannerPosX
// 		for _, char := range line {
// 			termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorGreen, termbox.ColorDefault)
// 			dynamicPosX++
// 		}
// 		dynamicPosY++
// 	}
// 	termbox.Sync()

// }
