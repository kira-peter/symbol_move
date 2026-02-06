package tetrisauto

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// TetrisAutoEffect 俄罗斯方块AI特效
type TetrisAutoEffect struct {
	tetris *TetrisAuto
	config *Config
}

// NewEffect 创建俄罗斯方块AI特效实例
func NewEffect() effects.Effect {
	return &TetrisAutoEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *TetrisAutoEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "tetris-auto",
		Name:        "俄罗斯方块AI",
		Description: "AI自动玩俄罗斯方块游戏",
		LongDescription: `
俄罗斯方块AI特效展示了AI自动玩经典俄罗斯方块游戏。

特点：
- 7种经典方块形状
- AI贪心策略（最低位置优先）
- 碰撞检测和消行动画
- 自动重启机制
- 彩色方块显示

完美用于：
- 游戏AI演示
- 经典游戏致敬
- 自动化展示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"游戏", "AI", "经典", "自动"},
	}
}

// Init 初始化特效
func (e *TetrisAutoEffect) Init(screen tcell.Screen) error {
	e.tetris = New(screen, e.config)
	return e.tetris.Init()
}

// Run 运行特效
func (e *TetrisAutoEffect) Run(quit <-chan struct{}) error {
	return e.tetris.Run(quit)
}

// Cleanup 清理资源
func (e *TetrisAutoEffect) Cleanup() error {
	if e.tetris != nil {
		return e.tetris.Cleanup()
	}
	return nil
}
