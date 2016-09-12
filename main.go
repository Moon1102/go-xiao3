package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {

	var g g2048

	err := termbox.Init()
	if err != nil {
		fmt.Printf("初始化失败, Error:%s\n", err)
		os.Exit(1)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputCurrent)

	g.Run()
}
