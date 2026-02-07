package rainbowwave

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// RainbowWaveEffect 彩虹波浪特效
type RainbowWaveEffect struct {
	wave   *RainbowWave
	config *Config
}

// NewEffect 创建彩虹波浪特效实例
func NewEffect() effects.Effect {
	return &RainbowWaveEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *RainbowWaveEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "rainbow-wave",
		Name:          "彩虹波浪",
		Description:   "彩虹色的水平波浪从左到右滚动,充满整个屏幕",
		NameEN:        "Rainbow Wave",
		DescriptionEN: "Rainbow-colored horizontal waves scrolling across the screen",
		LongDescription: `
彩虹波浪特效展示美丽的彩虹色波浪动画。

特点：
- 7种彩虹颜色循环
- 平滑波浪动画
- 多层波浪叠加
- 美观的 ASCII 字符
- 30 FPS 流畅运行
- 自动适配终端大小

完美用于：
- 放松心情的背景
- 色彩展示效果
- 装饰性动画
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"彩色", "波浪", "动画", "美观"},
	}
}

// Init 初始化特效
func (e *RainbowWaveEffect) Init(screen tcell.Screen) error {
	e.wave = New(screen, e.config)
	return e.wave.Init()
}

// Run 运行特效
func (e *RainbowWaveEffect) Run(quit <-chan struct{}) error {
	return e.wave.Run(quit)
}

// Cleanup 清理资源
func (e *RainbowWaveEffect) Cleanup() error {
	if e.wave != nil {
		return e.wave.Cleanup()
	}
	return nil
}
