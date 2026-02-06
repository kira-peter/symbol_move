package starrysky

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// StarrySkyEffect 星空闪烁特效
type StarrySkyEffect struct {
	sky    *StarrySky
	config *Config
}

// NewEffect 创建星空闪烁特效实例
func NewEffect() effects.Effect {
	return &StarrySkyEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *StarrySkyEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "starry-sky",
		Name:        "星空闪烁",
		Description: "模拟夜空中星星闪烁的效果",
		LongDescription: `
星空闪烁特效在终端中渲染一个美丽的夜空，数百颗星星随机闪烁。

特点：
- 真实的闪烁动画（基于正弦波）
- 多种颜色主题（经典、彩色、蓝色）
- 可调节星星密度
- 流畅的 30 FPS 动画
- 自动适配终端大小

完美用于：
- 放松心情的背景动画
- 演示终端渲染能力
- 作为其他特效的背景层
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"自然", "动画", "简单", "放松"},
	}
}

// Init 初始化特效
func (e *StarrySkyEffect) Init(screen tcell.Screen) error {
	e.sky = New(screen, e.config)
	return e.sky.Init()
}

// Run 运行特效
func (e *StarrySkyEffect) Run(quit <-chan struct{}) error {
	return e.sky.Run(quit)
}

// Cleanup 清理资源
func (e *StarrySkyEffect) Cleanup() error {
	if e.sky != nil {
		return e.sky.Cleanup()
	}
	return nil
}
