package main

//110 x 40 Terminal size
import (
	GoSnake "GoSnake/src"
	// "fmt"
	// "github.com/nsf/termbox-go"
)

func main() {

	// fmt.Println(GoSnake.MenuSelector(0))
	GoSnake.InitializeGame()
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
