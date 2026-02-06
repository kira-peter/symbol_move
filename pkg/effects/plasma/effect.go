package plasma

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type PlasmaEffect struct {
	plasma *Plasma
	config *Config
}

func NewEffect() effects.Effect {
	return &PlasmaEffect{
		config: DefaultConfig(),
	}
}

func (e *PlasmaEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "plasma",
		Name:        "Plasma 等离子",
		Description: "彩色等离子云效果，使用正弦函数生成图案，颜色循环动画",
		LongDescription: `
Plasma 等离子特效使用数学函数生成美丽的等离子云图案。

特点：
- 正弦波叠加算法
- 彩色渐变效果
- 平滑动画
- 数学艺术
- 30 FPS 流畅运行

完美用于：
- 数学可视化
- 科幻效果
- 背景动画
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"数学", "彩色", "等离子", "科幻"},
	}
}

func (e *PlasmaEffect) Init(screen tcell.Screen) error {
	e.plasma = New(screen, e.config)
	return e.plasma.Init()
}

func (e *PlasmaEffect) Run(quit <-chan struct{}) error {
	return e.plasma.Run(quit)
}

func (e *PlasmaEffect) Cleanup() error {
	if e.plasma != nil {
		return e.plasma.Cleanup()
	}
	return nil
}
