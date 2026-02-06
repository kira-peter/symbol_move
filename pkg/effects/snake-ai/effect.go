package snakeai

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// SnakeAIEffect 贪吃蛇AI特效
type SnakeAIEffect struct {
	snake  *SnakeAI
	config *Config
}

// NewEffect 创建贪吃蛇AI特效实例
func NewEffect() effects.Effect {
	return &SnakeAIEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *SnakeAIEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "snake-ai",
		Name:        "贪吃蛇AI",
		Description: "AI自动玩贪吃蛇游戏",
		LongDescription: `
贪吃蛇AI特效展示了AI自动玩经典贪吃蛇游戏。

特点：
- BFS路径规划算法
- 智能避免自撞
- 自动增长和食物生成
- 死亡自动重启
- 流畅的移动动画

完美用于：
- 游戏AI演示
- 路径规划算法展示
- 经典游戏致敬
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"游戏", "AI", "经典", "算法"},
	}
}

// Init 初始化特效
func (e *SnakeAIEffect) Init(screen tcell.Screen) error {
	e.snake = New(screen, e.config)
	return e.snake.Init()
}

// Run 运行特效
func (e *SnakeAIEffect) Run(quit <-chan struct{}) error {
	return e.snake.Run(quit)
}

// Cleanup 清理资源
func (e *SnakeAIEffect) Cleanup() error {
	if e.snake != nil {
		return e.snake.Cleanup()
	}
	return nil
}
