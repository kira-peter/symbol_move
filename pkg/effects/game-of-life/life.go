package gameoflife

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	InitDensity float64
	FPS         int
}

func DefaultConfig() *Config {
	return &Config{
		InitDensity: 0.3,
		FPS:         10,
	}
}

type GameOfLife struct {
	screen  tcell.Screen
	config  *Config
	grid    [][]bool
	newGrid [][]bool
	width   int
	height  int
	rand    *rand.Rand
}

func New(screen tcell.Screen, config *Config) *GameOfLife {
	if config == nil {
		config = DefaultConfig()
	}

	return &GameOfLife{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *GameOfLife) Init() error {
	g.width, g.height = g.screen.Size()
	g.grid = make([][]bool, g.height)
	g.newGrid = make([][]bool, g.height)

	for y := 0; y < g.height; y++ {
		g.grid[y] = make([]bool, g.width)
		g.newGrid[y] = make([]bool, g.width)

		for x := 0; x < g.width; x++ {
			g.grid[y][x] = g.rand.Float64() < g.config.InitDensity
		}
	}

	return nil
}

func (g *GameOfLife) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := (x + dx + g.width) % g.width
			ny := (y + dy + g.height) % g.height

			if g.grid[ny][nx] {
				count++
			}
		}
	}
	return count
}

func (g *GameOfLife) Update() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			neighbors := g.countNeighbors(x, y)
			alive := g.grid[y][x]

			// Conway's Game of Life 规则
			if alive {
				g.newGrid[y][x] = neighbors == 2 || neighbors == 3
			} else {
				g.newGrid[y][x] = neighbors == 3
			}
		}
	}

	// 交换缓冲区
	g.grid, g.newGrid = g.newGrid, g.grid
}

func (g *GameOfLife) Render() {
	g.screen.Clear()

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.grid[y][x] {
				g.screen.SetContent(x, y, '●', nil, style)
			}
		}
	}

	g.screen.Show()
}

func (g *GameOfLife) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(g.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			g.Update()
			g.Render()
		}
	}
}

func (g *GameOfLife) Cleanup() error {
	return nil
}
