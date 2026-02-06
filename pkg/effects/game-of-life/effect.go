package gameoflife

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type GameOfLifeEffect struct {
	game   *GameOfLife
	config *Config
}

func NewEffect() effects.Effect {
	return &GameOfLifeEffect{
		config: DefaultConfig(),
	}
}

func (e *GameOfLifeEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "game-of-life",
		Name:        "生命游戏",
		Description: "Conway's Game of Life细胞自动机,随机初始状态",
		LongDescription: `
生命游戏特效实现了经典的 Conway's Game of Life 算法。

特点：
- 经典细胞自动机
- 随机初始状态
- 自动演化
- 绿色细胞显示
- 支持边界循环

完美用于：
- 算法演示
- 数学教学
- 自动机展示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"算法", "自动机", "经典", "模拟"},
	}
}

func (e *GameOfLifeEffect) Init(screen tcell.Screen) error {
	e.game = New(screen, e.config)
	return e.game.Init()
}

func (e *GameOfLifeEffect) Run(quit <-chan struct{}) error {
	return e.game.Run(quit)
}

func (e *GameOfLifeEffect) Cleanup() error {
	if e.game != nil {
		return e.game.Cleanup()
	}
	return nil
}
