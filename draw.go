package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func DrawSurface() {
	// 绘制边框
	termbox.SetCell(20, 0, 0x250C, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(60, 0, 0x2510, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(20, 24, 0x2514, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(60, 24, 0x2518, termbox.ColorYellow, termbox.ColorDefault)

	for i := 21; i < 60; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 6, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 12, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 18, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 24, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
	}

	for i := 1; i < 24; i++ {
		termbox.SetCell(20, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(30, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(40, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(50, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(60, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
	}
}

func DrawElement(a *array) {
	for i, tn := range a {
		for j, v := range tn {
			if v != 0 {
				printElement(a, elts[v], i, j)
			}
		}
	}
}

func printElement(a *array, e []attr, row, col int) {
	for _, k := range e {
		termbox.SetCell((col+2)*10+k.x, row*6+k.y, k.ch, k.fg, k.bg)
	}

}

// 刷新
func Flush() {
	termbox.Flush()
}

// 清屏
func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Draw(a *array) {
	Clear()
	DrawSurface()
	DrawElement(a)
	Flush()
}

func PrintGameover(array [4][4]int) {
	Clear()

	var total int

	for _, temp := range array {
		for _, score := range temp {
			total += score
		}
	}

	for i, v := range "Total: " + fmt.Sprintf("%d", total) {
		termbox.SetCell(35+i, 10, v, termbox.ColorYellow, termbox.ColorDefault)
	}

	for i, v := range "GAMEOVER" {
		termbox.SetCell(36+i, 11, v, termbox.ColorYellow, termbox.ColorDefault)
	}

	commonPrint()

	Flush()
}

func PrintWin() {
	Clear()

	for i, v := range "WINNING" {
		termbox.SetCell(36+i, 11, v, termbox.ColorYellow, termbox.ColorDefault)
	}

	commonPrint()

	Flush()
}

func commonPrint() {

	for i, v := range "Ctrl+R: Continue" {
		termbox.SetCell(32+i, 12, v, termbox.ColorYellow, termbox.ColorDefault)
	}

	for i, v := range "Ctrl+Q: Exit" {
		termbox.SetCell(34+i, 13, v, termbox.ColorYellow, termbox.ColorDefault)
	}
}
