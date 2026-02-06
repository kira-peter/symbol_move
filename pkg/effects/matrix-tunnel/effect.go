package matrixtunnel

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// MatrixTunnelEffect 矩阵隧道特效
type MatrixTunnelEffect struct {
	tunnel *MatrixTunnel
	config *Config
}

// NewEffect 创建矩阵隧道特效实例
func NewEffect() effects.Effect {
	return &MatrixTunnelEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *MatrixTunnelEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "matrix-tunnel",
		Name:        "矩阵隧道",
		Description: "模拟飞行穿越矩阵隧道的3D效果",
		LongDescription: `
矩阵隧道特效展示了飞行穿越充满矩阵字符的3D隧道。

特点：
- 3D透视投影
- 深度缓冲和字符亮度
- 流畅的飞行动画
- 无限循环深度
- 绿色矩阵主题

完美用于：
- 科幻主题展示
- 黑客风格背景
- 3D效果演示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"3D", "科幻", "动画", "矩阵"},
	}
}

// Init 初始化特效
func (e *MatrixTunnelEffect) Init(screen tcell.Screen) error {
	e.tunnel = New(screen, e.config)
	return e.tunnel.Init()
}

// Run 运行特效
func (e *MatrixTunnelEffect) Run(quit <-chan struct{}) error {
	return e.tunnel.Run(quit)
}

// Cleanup 清理资源
func (e *MatrixTunnelEffect) Cleanup() error {
	if e.tunnel != nil {
		return e.tunnel.Cleanup()
	}
	return nil
}
