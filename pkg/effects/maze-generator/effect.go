package mazegenerator

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type MazeGeneratorEffect struct {
	maze   *MazeGenerator
	config *Config
}

func NewEffect() effects.Effect {
	return &MazeGeneratorEffect{
		config: DefaultConfig(),
	}
}

func (e *MazeGeneratorEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "maze-generator",
		Name:        "迷宫生成",
		Description: "实时展示迷宫生成过程,使用DFS算法",
		LongDescription: `
迷宫生成特效展示深度优先搜索算法生成迷宫的过程。

特点：
- DFS 算法实时演示
- 彩色路径显示
- 平滑生成动画
- 完成后保持显示

完美用于：
- 算法教学
- 迷宫展示
- 路径规划演示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"算法", "迷宫", "生成", "路径"},
	}
}

func (e *MazeGeneratorEffect) Init(screen tcell.Screen) error {
	e.maze = New(screen, e.config)
	return e.maze.Init()
}

func (e *MazeGeneratorEffect) Run(quit <-chan struct{}) error {
	return e.maze.Run(quit)
}

func (e *MazeGeneratorEffect) Cleanup() error {
	if e.maze != nil {
		return e.maze.Cleanup()
	}
	return nil
}
