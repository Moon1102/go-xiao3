package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type g2048 struct {
	a          array
	width      int
	height     int
	winScore   int  // 游戏胜利的分数
	totalScore int  //总分
	nelts      int  // 当前的元素的个数
	moved      bool // 是否可以移动
	eq         bool // 是否右相等的元素
	exit       bool // 退出程序
	win        bool // 游戏成功
	gameover   bool // 游戏失败
	pause      bool // 等待重新开始游戏

}

type EmptyElem struct {
	row, col int
}

type array [4][4]int

const (
	UpOP = iota
	RightOP
	DownOP
	LeftOP
)

// 初始化
func (t *g2048) Init() {
	t.width = 4
	t.height = 4
	t.winScore = 768
	t.nelts = 0
	t.a = array{}

	t.eq = true
	t.moved = true
	t.win = false
	t.gameover = false
	t.exit = false
	t.pause = false

	t.Rand()
	Draw(&t.a)
}

// 循环监听事件
func (t *g2048) loop() {
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			switch ev.Key {
			case termbox.KeyCtrlR:
				break loop
			case termbox.KeyCtrlQ:
				t.exit = true
				break loop

			case termbox.KeyArrowLeft:
				t.handle(LeftOP)

			case termbox.KeyArrowUp:
				t.handle(UpOP)

			case termbox.KeyArrowRight:
				t.handle(RightOP)

			case termbox.KeyArrowDown:
				t.handle(DownOP)
			}
		}
	}
}

// 运行程序
func (t *g2048) Run() {
	for {
		t.Init()

		t.loop()

		if t.exit {
			return
		}
	}
}

// 根据不同的操作进行处理
func (t *g2048) handle(op int) {

	if !t.IsPause() {
		switch op {
		case UpOP:
			t.MoveUp()
		case RightOP:
			t.MoveRight()
		case DownOP:
			t.MoveDown()
		case LeftOP:
			t.MoveLeft()
		}

		if !t.IsWin() {
			t.Rand()

			if !t.IsGameover() {
				Draw(&t.a)
			}
		}
	}
}

// 向上移动
func (t *g2048) MoveUp() {
	t.moved = false
	t.eq = false

	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			t.Merge(i, j)
		}
	}
}

// 向下移动
func (t *g2048) MoveDown() {
	t.Rollback()

	t.MoveUp()

	t.Rollback()
}

// 向左移动
func (t *g2048) MoveLeft() {
	t.TurnRight()

	t.MoveUp()

	t.TurnLeft()
}

// 向右移动
func (t *g2048) MoveRight() {
	t.TurnLeft()

	t.MoveUp()

	t.TurnRight()
}

// 旋转180度
func (t *g2048) Rollback() {
	for i := 0; i < t.height/2; i++ {
		for j := 0; j < t.width; j++ {
			t.a[i][j], t.a[t.height-i-1][t.width-j-1] =
				t.a[t.height-i-1][t.width-j-1], t.a[i][j]
		}
	}
}

// 向右旋转90度
func (t *g2048) TurnRight() {
	for i := 0; i < t.height-1; i++ {
		for j := i + 1; j < t.width; j++ {
			t.a[i][j], t.a[j][i] = t.a[j][i], t.a[i][j]
		}
	}

	for i := 0; i < t.height/2; i++ {
		for j := 0; j < t.width; j++ {
			t.a[j][i], t.a[j][t.height-i-1] = t.a[j][t.height-i-1], t.a[j][i]
		}
	}
}

// 向左旋转90度
func (t *g2048) TurnLeft() {
	for i := 0; i < t.height/2; i++ {
		for j := 0; j < t.width; j++ {
			t.a[j][i], t.a[j][t.height-i-1] = t.a[j][t.height-i-1], t.a[j][i]
		}
	}

	for i := 0; i < t.height-1; i++ {
		for j := i + 1; j < t.width; j++ {
			t.a[i][j], t.a[j][i] = t.a[j][i], t.a[i][j]
		}
	}
}

// 合并当前元素的下一个元素
func (t *g2048) Merge(i, j int) {
	t.move(i, j)

	if t.a[i][j] != 0 && i+1 < t.height {

		if t.a[i][j] == t.a[i+1][j] && t.a[i][j] != 1 && t.a[i][j] != 2 { // 当前元素和下一个元素相等,且非1，非2

			t.a[i][j], t.a[i+1][j] = t.a[i][j]+t.a[i+1][j], 0
			t.nelts--

			if t.a[i][j] == t.winScore {
				t.win = true
			}

		} else {
			if t.a[i][j] == 1 && t.a[i+1][j] == 2 || t.a[i][j] == 2 && t.a[i+1][j] == 1 {
				t.a[i][j], t.a[i+1][j] = t.a[i][j]+t.a[i+1][j], 0
				t.nelts--
			}

			if t.a[i][j] == t.winScore {
				t.win = true
			}
		}
	}
}

// 将第i列第j行后面的所有非0值向上移动
func (t *g2048) move(i, j int) {

	for v := i; v < t.height; v++ {

		if t.a[v][j] == 0 {

			for k := v + 1; k < t.height; k++ {

				if t.a[k][j] != 0 { // 向上移动非0的元素
					t.a[v][j], t.a[k][j] = t.a[k][j], 0
					t.moved = true
					break
				}
			}
		}
	}

}

// 在为0值的方格内随机生成
func (t *g2048) Rand() {
	var list []EmptyElem
	var val EmptyElem

	if !t.moved {
		return
	}

	// 如果元素满了，且每个元素都不能进行合并，游戏结束
	if t.nelts++; t.nelts == t.height*t.width && !t.IsMerge() {
		t.gameover = true
		return
	}

	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if t.a[i][j] == 0 {
				val.row, val.col = i, j
				list = append(list, val) // 将为0的元素加入到切片中
			}
		}
	}

	// 在元素为0的切片中随机选择一个位置，在此位置上随机产生一个0或4的数
	v := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(list))
	//产生一个0-100的数
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)

	//一开始不要产生很大的数
	if len(list) == 16 {
		t.a[list[v].row][list[v].col] = 1
		return
	}
	if len(list) == 15 {
		t.a[list[v].row][list[v].col] = 2
		return
	}

	//正常随机
	switch r % 10 {
	case 0, 1:
		t.a[list[v].row][list[v].col] = 1
	case 2, 3:
		t.a[list[v].row][list[v].col] = 2
	case 4, 5, 6, 7:
		t.a[list[v].row][list[v].col] = 3
	case 8:
		t.a[list[v].row][list[v].col] = 6
	//此时值为9 + (r/10 * n); n = [0,9]
	default:
		//所以对值做限定，保证不要一下出很大的值
		switch r / 10 {
		case 0, 1, 6, 3, 8:
			t.a[list[v].row][list[v].col] = 12

		case 5, 2, 7:
			t.a[list[v].row][list[v].col] = 24

		case 4, 9:
			t.a[list[v].row][list[v].col] = 48
		}

	}
}

// 检查是否可以合并元素
func (t *g2048) IsMerge() bool {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {

			if i+1 < t.height {
				switch t.a[i][j] {
				case 1:
					if t.a[i+1][j] == 2 {
						return true
					}
				case 2:
					if t.a[i+1][j] == 1 {
						return true
					}
				default:
					if t.a[i][j] == t.a[i+1][j] && t.a[i][j] > 2 {
						return true
					} else {
						return false
					}
				}
			}

			if j+1 < t.width {
				switch t.a[i][j] {
				case 1:
					if t.a[i][j+1] == 2 {
						return true
					}
				case 2:
					if t.a[i][j+1] == 1 {
						return true
					}
				default:
					if t.a[i][j] == t.a[i][j+1] && t.a[i][j] > 2 {
						return true
					} else {
						return false
					}
				}
			}

		}
	}

	return false
}

// 游戏是否成功
func (t *g2048) IsWin() bool {
	if t.win {
		t.pause = true

		PrintWin()
		return true
	}

	return false
}

// 游戏是否失败
func (t *g2048) IsGameover() bool {
	if t.gameover {
		t.pause = true

		PrintGameover(t.a)
		return true
	}

	return false
}

// 游戏是否暂停
func (t *g2048) IsPause() bool {
	if t.pause {
		return true
	}

	return false
}
