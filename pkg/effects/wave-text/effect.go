package wavetext

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// WaveTextEffect 波浪文字特效
type WaveTextEffect struct {
	wave   *WaveText
	config *Config
}

// NewEffect 创建波浪文字特效实例
func NewEffect() effects.Effect {
	return &WaveTextEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *WaveTextEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "wave-text",
		Name:          "波浪文字",
		Description:   "显示文字以正弦波形式上下波动,配有彩色渐变效果",
		NameEN:        "Wave Text",
		DescriptionEN: "Text displayed in a sine wave pattern with rainbow gradient",
		LongDescription: `
波浪文字特效让文本像波浪一样上下起伏，并伴随彩虹渐变。

特点：
- 正弦波平滑动画
- 彩虹渐变色彩
- 流畅的波浪效果
- 居中显示
- 30 FPS 流畅运行
- 自动适配终端大小

完美用于：
- 欢迎界面动画
- Logo 展示效果
- 文字特效演示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"文字", "动画", "彩色", "流畅"},
	}
}

// Init 初始化特效
func (e *WaveTextEffect) Init(screen tcell.Screen) error {
	e.wave = New(screen, e.config)
	return e.wave.Init()
}

// Run 运行特效
func (e *WaveTextEffect) Run(quit <-chan struct{}) error {
	return e.wave.Run(quit)
}

// Cleanup 清理资源
func (e *WaveTextEffect) Cleanup() error {
	if e.wave != nil {
		return e.wave.Cleanup()
	}
	return nil
}
