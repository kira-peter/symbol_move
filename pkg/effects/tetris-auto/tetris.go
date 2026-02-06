package tetrisauto

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 俄罗斯方块配置
type Config struct {
	FallSpeed float64 // 下落速度（行/秒）
	FPS       int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		FallSpeed: 2.0, // 2行/秒
		FPS:       30,
	}
}

// Tetromino 方块
type Tetromino struct {
	shape [][]int
	x, y  int
	color tcell.Color
}

// tetrominoes 7种经典方块形状
var tetrominoes = [][][]int{
	{{1, 1, 1, 1}},                   // I
	{{1, 1}, {1, 1}},                 // O
	{{0, 1, 0}, {1, 1, 1}},           // T
	{{0, 1, 1}, {1, 1, 0}},           // S
	{{1, 1, 0}, {0, 1, 1}},           // Z
	{{1, 0, 0}, {1, 1, 1}},           // J
	{{0, 0, 1}, {1, 1, 1}},           // L
}

var tetrominoColors = []tcell.Color{
	tcell.ColorLightCyan, // I
	tcell.ColorYellow,    // O
	tcell.ColorPurple,    // T
	tcell.ColorGreen,     // S
	tcell.ColorRed,       // Z
	tcell.ColorBlue,      // J
	tcell.ColorOrange,    // L
}

// TetrisAuto 俄罗斯方块AI特效
type TetrisAuto struct {
	screen     tcell.Screen
	config     *Config
	board      [][]int // 游戏板（0=空，1-7=不同方块颜色）
	boardColors [][]tcell.Color
	current    *Tetromino
	fallTimer  float64
	width      int
	height     int
	boardW     int
	boardH     int
	score      int
	lastUpdate time.Time
	rand       *rand.Rand
}

// New 创建俄罗斯方块AI特效实例
func New(screen tcell.Screen, config *Config) *TetrisAuto {
	if config == nil {
		config = DefaultConfig()
	}

	return &TetrisAuto{
		screen: screen,
		config: config,
		boardW: 10,
		boardH: 20,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化俄罗斯方块
func (t *TetrisAuto) Init() error {
	t.width, t.height = t.screen.Size()
	t.board = make([][]int, t.boardH)
	t.boardColors = make([][]tcell.Color, t.boardH)
	for i := range t.board {
		t.board[i] = make([]int, t.boardW)
		t.boardColors[i] = make([]tcell.Color, t.boardW)
	}
	t.fallTimer = 0
	t.score = 0
	t.spawnNew()
	t.lastUpdate = time.Now()
	return nil
}

// spawnNew 生成新方块
func (t *TetrisAuto) spawnNew() {
	shapeIdx := t.rand.Intn(len(tetrominoes))
	t.current = &Tetromino{
		shape: tetrominoes[shapeIdx],
		x:     t.boardW/2 - len(tetrominoes[shapeIdx][0])/2,
		y:     0,
		color: tetrominoColors[shapeIdx],
	}
}

// canMove 检测是否可以移动到指定位置
func (t *TetrisAuto) canMove(x, y int, shape [][]int) bool {
	for dy, row := range shape {
		for dx, cell := range row {
			if cell == 0 {
				continue
			}

			newX := x + dx
			newY := y + dy

			// 边界检测
			if newX < 0 || newX >= t.boardW || newY >= t.boardH {
				return false
			}

			// 碰撞检测（跳过负数y，允许从顶部进入）
			if newY >= 0 && t.board[newY][newX] != 0 {
				return false
			}
		}
	}
	return true
}

// placeTetromino 放置方块到游戏板
func (t *TetrisAuto) placeTetromino() {
	for dy, row := range t.current.shape {
		for dx, cell := range row {
			if cell == 0 {
				continue
			}

			newY := t.current.y + dy
			newX := t.current.x + dx

			if newY >= 0 && newY < t.boardH && newX >= 0 && newX < t.boardW {
				t.board[newY][newX] = 1
				t.boardColors[newY][newX] = t.current.color
			}
		}
	}
}

// clearLines 消除完整的行
func (t *TetrisAuto) clearLines() {
	linesCleared := 0
	for y := t.boardH - 1; y >= 0; y-- {
		full := true
		for x := 0; x < t.boardW; x++ {
			if t.board[y][x] == 0 {
				full = false
				break
			}
		}

		if full {
			linesCleared++
			// 删除这一行，上面的行下移
			for y2 := y; y2 > 0; y2-- {
				t.board[y2] = t.board[y2-1]
				t.boardColors[y2] = t.boardColors[y2-1]
			}
			t.board[0] = make([]int, t.boardW)
			t.boardColors[0] = make([]tcell.Color, t.boardW)
			y++ // 重新检查当前行
		}
	}
	t.score += linesCleared
}

// findBestMove AI寻找最佳移动
func (t *TetrisAuto) findBestMove() int {
	bestScore := -99999
	bestX := t.current.x

	// 简单AI：找到最低的位置
	for x := 0; x < t.boardW; x++ {
		// 测试这个x位置能下降到哪里
		testY := t.current.y
		for testY < t.boardH && t.canMove(x, testY+1, t.current.shape) {
			testY++
		}

		// 评分：越低越好，但要避免制造空洞
		score := testY * 10

		// 检查是否会造成空洞
		holes := 0
		for dy, row := range t.current.shape {
			for dx, cell := range row {
				if cell == 0 {
					continue
				}
				// 边界检查
				checkX := x + dx
				if checkX < 0 || checkX >= t.boardW {
					continue
				}
				// 检查下方是否有空洞
				checkY := testY + dy + 1
				if checkY < t.boardH {
					for cy := checkY; cy < t.boardH; cy++ {
						if t.board[cy][checkX] == 0 {
							holes++
						}
					}
				}
			}
		}
		score -= holes * 20

		if score > bestScore {
			bestScore = score
			bestX = x
		}
	}

	return bestX
}

// Update 更新俄罗斯方块状态
func (t *TetrisAuto) Update(deltaTime float64) {
	t.fallTimer += deltaTime

	if t.fallTimer >= 1.0/t.config.FallSpeed {
		t.fallTimer = 0

		// AI决策：找最佳位置
		bestX := t.findBestMove()
		t.current.x = bestX

		// 下落
		if t.canMove(t.current.x, t.current.y+1, t.current.shape) {
			t.current.y++
		} else {
			// 放置方块
			t.placeTetromino()
			t.clearLines()
			t.spawnNew()

			// 检测游戏结束
			if !t.canMove(t.current.x, t.current.y, t.current.shape) {
				// 重新开始
				for i := range t.board {
					for j := range t.board[i] {
						t.board[i][j] = 0
					}
				}
				t.score = 0
			}
		}
	}
}

// Render 渲染俄罗斯方块
func (t *TetrisAuto) Render() {
	t.screen.Clear()

	// 计算居中位置
	offsetX := (t.width - t.boardW*2) / 2
	offsetY := (t.height - t.boardH) / 2

	// 绘制边框
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	for y := 0; y <= t.boardH; y++ {
		t.screen.SetContent(offsetX-1, offsetY+y, '│', nil, borderStyle)
		t.screen.SetContent(offsetX+t.boardW*2, offsetY+y, '│', nil, borderStyle)
	}
	for x := 0; x <= t.boardW*2; x++ {
		t.screen.SetContent(offsetX+x-1, offsetY+t.boardH, '─', nil, borderStyle)
	}

	// 绘制游戏板
	for y := 0; y < t.boardH; y++ {
		for x := 0; x < t.boardW; x++ {
			if t.board[y][x] != 0 {
				style := tcell.StyleDefault.Foreground(t.boardColors[y][x]).Bold(true)
				t.screen.SetContent(offsetX+x*2, offsetY+y, '█', nil, style)
				t.screen.SetContent(offsetX+x*2+1, offsetY+y, '█', nil, style)
			}
		}
	}

	// 绘制当前方块
	if t.current != nil {
		for dy, row := range t.current.shape {
			for dx, cell := range row {
				if cell == 0 {
					continue
				}

				screenY := offsetY + t.current.y + dy
				screenX := offsetX + (t.current.x+dx)*2

				if screenY >= offsetY && screenY < offsetY+t.boardH {
					style := tcell.StyleDefault.Foreground(t.current.color).Bold(true)
					t.screen.SetContent(screenX, screenY, '█', nil, style)
					t.screen.SetContent(screenX+1, screenY, '█', nil, style)
				}
			}
		}
	}

	t.screen.Show()
}

// Run 运行俄罗斯方块特效
func (t *TetrisAuto) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(t.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(t.lastUpdate).Seconds()
			t.lastUpdate = now

			t.Update(deltaTime)
			t.Render()
		}
	}
}

// Cleanup 清理资源
func (t *TetrisAuto) Cleanup() error {
	return nil
}
