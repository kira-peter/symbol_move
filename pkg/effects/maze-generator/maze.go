package mazegenerator

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	CellSize int
	Speed    int
	FPS      int
}

func DefaultConfig() *Config {
	return &Config{
		CellSize: 2,
		Speed:    5,
		FPS:      30,
	}
}

type Cell struct {
	x, y    int
	visited bool
	walls   [4]bool // 上右下左
}

type MazeGenerator struct {
	screen  tcell.Screen
	config  *Config
	cells   [][]*Cell
	width   int
	height  int
	rows    int
	cols    int
	stack   []*Cell
	current *Cell
	done    bool
	rand    *rand.Rand
}

func New(screen tcell.Screen, config *Config) *MazeGenerator {
	if config == nil {
		config = DefaultConfig()
	}

	return &MazeGenerator{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m *MazeGenerator) Init() error {
	m.width, m.height = m.screen.Size()
	m.cols = m.width / m.config.CellSize
	m.rows = m.height / m.config.CellSize

	// 初始化迷宫
	m.cells = make([][]*Cell, m.rows)
	for y := 0; y < m.rows; y++ {
		m.cells[y] = make([]*Cell, m.cols)
		for x := 0; x < m.cols; x++ {
			m.cells[y][x] = &Cell{
				x:       x,
				y:       y,
				visited: false,
				walls:   [4]bool{true, true, true, true},
			}
		}
	}

	m.current = m.cells[0][0]
	m.current.visited = true
	m.stack = []*Cell{}
	m.done = false

	return nil
}

func (m *MazeGenerator) getUnvisitedNeighbor(cell *Cell) *Cell {
	neighbors := []*Cell{}

	// 上
	if cell.y > 0 && !m.cells[cell.y-1][cell.x].visited {
		neighbors = append(neighbors, m.cells[cell.y-1][cell.x])
	}
	// 右
	if cell.x < m.cols-1 && !m.cells[cell.y][cell.x+1].visited {
		neighbors = append(neighbors, m.cells[cell.y][cell.x+1])
	}
	// 下
	if cell.y < m.rows-1 && !m.cells[cell.y+1][cell.x].visited {
		neighbors = append(neighbors, m.cells[cell.y+1][cell.x])
	}
	// 左
	if cell.x > 0 && !m.cells[cell.y][cell.x-1].visited {
		neighbors = append(neighbors, m.cells[cell.y][cell.x-1])
	}

	if len(neighbors) > 0 {
		return neighbors[m.rand.Intn(len(neighbors))]
	}
	return nil
}

func (m *MazeGenerator) removeWalls(current, next *Cell) {
	dx := next.x - current.x
	dy := next.y - current.y

	if dx == 1 {
		current.walls[1] = false // 右
		next.walls[3] = false    // 左
	} else if dx == -1 {
		current.walls[3] = false // 左
		next.walls[1] = false    // 右
	} else if dy == 1 {
		current.walls[2] = false // 下
		next.walls[0] = false    // 上
	} else if dy == -1 {
		current.walls[0] = false // 上
		next.walls[2] = false    // 下
	}
}

func (m *MazeGenerator) Update() {
	if m.done {
		return
	}

	for i := 0; i < m.config.Speed && !m.done; i++ {
		next := m.getUnvisitedNeighbor(m.current)

		if next != nil {
			next.visited = true
			m.stack = append(m.stack, m.current)
			m.removeWalls(m.current, next)
			m.current = next
		} else if len(m.stack) > 0 {
			m.current = m.stack[len(m.stack)-1]
			m.stack = m.stack[:len(m.stack)-1]
		} else {
			m.done = true
		}
	}
}

func (m *MazeGenerator) Render() {
	m.screen.Clear()

	for y := 0; y < m.rows; y++ {
		for x := 0; x < m.cols; x++ {
			cell := m.cells[y][x]
			m.drawCell(cell)
		}
	}

	m.screen.Show()
}

func (m *MazeGenerator) drawCell(cell *Cell) {
	sx := cell.x * m.config.CellSize
	sy := cell.y * m.config.CellSize

	var style tcell.Style
	if cell == m.current {
		style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	} else if cell.visited {
		style = tcell.StyleDefault.Foreground(tcell.ColorGreen)
	} else {
		style = tcell.StyleDefault.Foreground(tcell.ColorGray)
	}

	// 绘制墙壁
	if cell.walls[0] && sy > 0 { // 上
		for i := 0; i < m.config.CellSize; i++ {
			if sx+i < m.width {
				m.screen.SetContent(sx+i, sy, '─', nil, style)
			}
		}
	}
	if cell.walls[1] && sx+m.config.CellSize-1 < m.width { // 右
		for i := 0; i < m.config.CellSize; i++ {
			if sy+i < m.height {
				m.screen.SetContent(sx+m.config.CellSize-1, sy+i, '│', nil, style)
			}
		}
	}
}

func (m *MazeGenerator) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(m.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			m.Update()
			m.Render()
		}
	}
}

func (m *MazeGenerator) Cleanup() error {
	return nil
}
