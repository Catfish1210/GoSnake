package GoSnake

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Options struct {
	display    [][]string
	active     int
	banner     []string
	difficulty int
}

type MenuState struct {
	active     int
	difficulty int
}

// write difficulty to this struct and to this function below me, so you can still use the recursion but also return
// the difficulty once game option 0 (play) is selected
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
					//
					InitializeGame()
					//
					return menuOptions.active
				} else if menuOptions.active == 1 {
					menuOptions.difficulty = getDifficulty(menuOptions)
					menuOptions.banner = Banner
					_ = generateBanner(menuOptions)
					fmt.Println(menuOptions.difficulty)
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

func getDifficulty(menuOptions Options) int {

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

	menuOptions.banner = BannerEmpty
	_ = generateBanner(menuOptions)

	terminalWidth, terminalHeight := termbox.Size()
	difficultyPosY := (terminalHeight / 8) + 1
	difficultyPosX := (terminalWidth / 2) - (len(Banner[1]) / 2) + 9
	difficultyCords := []int{difficultyPosX, difficultyPosY}

	active := 0
	updateDifficultyDisplay(difficultyCords, active)
	termbox.Sync()

	returnVal := -1
	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyCtrlC {
				return returnVal
			}
			if keySeq.Key == termbox.KeyEsc {
				MenuSelector(1)
			}
			if keySeq.Ch == 'w' {
				if active > 0 {
					active -= 1
					updateDifficultyDisplay(difficultyCords, active)
				}
			} else if keySeq.Ch == 's' {
				if active < 2 {
					active += 1
					updateDifficultyDisplay(difficultyCords, active)
				}
			} else if keySeq.Key == termbox.KeyEnter || keySeq.Ch == 'd' {
				go updateDifficultySet(difficultyCords, active)
				if menuOptions.active == 0 {
					returnVal = 0
				} else if menuOptions.active == 1 {
					returnVal = 1
				} else if menuOptions.active == 2 {
					returnVal = 2
				}
			}
		}
	}

}
func updateDifficultySet(difficultyCords []int, active int) {
	termbox.SetCell(difficultyCords[0], difficultyCords[1]+4, 'X', termbox.ColorDefault, termbox.ColorDefault)

	initialX, initialY := difficultyCords[0], difficultyCords[1]+4

	for i := 0; i < 33; i++ {
		termbox.SetCell(initialX+i, initialY, ' ', termbox.ColorDefault, termbox.ColorDefault)
		termbox.Sync()
	}

	difOptions := []string{" Game Level Set To:    EASY", " Game Level Set To:    MODERATE", " Game Level Set To:  HARD"}
	for i, option := range difOptions {
		if i == active {
			dynamicX := initialX
			for _, char := range option {
				termbox.SetCell(dynamicX, initialY, char, termbox.ColorRed, termbox.ColorDefault)
				dynamicX++
			}
		}
	}
	termbox.Sync()
	time.Sleep(3 * time.Second)
	for i := 0; i < 33; i++ {
		termbox.SetCell(initialX+i, initialY, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Sync()

}

func updateDifficultyDisplay(difficultyCords []int, active int) {
	difOptions := []string{"1    ---    Easy", "2    ---    Moderate", "3    ---    Hard"}

	initialX, initialY := difficultyCords[0], difficultyCords[1]

	for i, line := range difOptions {
		dynamicX := initialX
		for _, char := range line {
			if i == active {
				termbox.SetCell(dynamicX, initialY, char, termbox.ColorRed, termbox.ColorDefault)
			} else {
				termbox.SetCell(dynamicX, initialY, char, termbox.ColorDefault, termbox.ColorDefault)
			}
			dynamicX++
		}
		initialY++
	}
	termbox.Sync()
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
