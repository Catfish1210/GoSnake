package GoSnake

import (
	"os"

	"github.com/nsf/termbox-go"
)

type Options struct {
	display [][]string
	active  int
	banner  []string
}

func MenuSelector(preselect int) int {
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

	var menuOptions Options
	menuOptions.display = append(menuOptions.display, PlayOption, LevelOption, HighscoresOption, QuitOption)
	menuOptions.active = preselect
	menuOptions.banner = Banner

	dynamicPosY := generateBanner(menuOptions)
	updateMenuDisplay(menuOptions, dynamicPosY)

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				termbox.Close()
				break
			}

			if keySeq.Ch == 'w' {
				if menuOptions.active > 0 {
					menuOptions.active -= 1
					updateMenuDisplay(menuOptions, dynamicPosY)
				}

			} else if keySeq.Ch == 's' {
				if menuOptions.active < 3 {
					menuOptions.active += 1
					updateMenuDisplay(menuOptions, dynamicPosY)
				}
			} else if keySeq.Key == termbox.KeyEnter || keySeq.Ch == 'd' {
				if menuOptions.active == 0 {
					return menuOptions.active
				} else if menuOptions.active == 1 {
					//inner menu (another menu inside of banner)
					menuOptions.banner = BannerEmpty
					_ = generateBanner(menuOptions)
					termbox.Sync()
				} else if menuOptions.active == 2 {
					//display scores in empty banner(include back function to menu)
				} else if menuOptions.active == 3 {
					termbox.Close()
					os.Exit(0)
				}

			}
		}
	}
	return -1
}

func generateBanner(menuOptions Options) int {
	terminalWidth, terminalHeight := termbox.Size()
	bannerPosY := (terminalHeight / 8) - 2
	bannerPosX := (terminalWidth / 2) - (len(Banner[1]) / 2)
	dynamicPosY := bannerPosY
	for _, line := range menuOptions.banner {
		dynamicPosX := bannerPosX
		for _, char := range line {
			termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorGreen, termbox.ColorDefault)
			dynamicPosX++
		}
		dynamicPosY++
	}
	return dynamicPosY
}

func updateMenuDisplay(menuOptions Options, dynamicPosY int) {
	terminalWidth, _ := termbox.Size()
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
				if i == menuOptions.active {
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
	termbox.Sync()
}
