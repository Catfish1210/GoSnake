package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	fmt.Println("Test")
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

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Ch == 'h' {
				fmt.Println("Hello World!")
			} else if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				break
			}
		}
	}

}
