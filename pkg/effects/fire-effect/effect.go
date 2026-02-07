package fireeffect

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type FireEffectEffect struct {
	fire   *FireEffect
	config *Config
}

func NewEffect() effects.Effect {
	return &FireEffectEffect{
		config: DefaultConfig(),
	}
}

func (e *FireEffectEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "fire-effect",
		Name:          "火焰燃烧",
		Description:   "从底部向上燃烧的火焰效果,使用字符密度模拟火焰形状",
		NameEN:        "Fire Effect",
		DescriptionEN: "Burning fire effect rising from bottom, using character density to simulate flames",
		LongDescription: `
火焰燃烧特效模拟真实的火焰向上燃烧效果。

特点：
- 真实的火焰模拟算法
- 红黄渐变色
- 热量传播效果
- 字符密度变化
- 30 FPS 流畅运行

完美用于：
- 火焰效果展示
- 热力学模拟
- 视觉特效演示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"火焰", "热", "动画", "自然"},
	}
}

func (e *FireEffectEffect) Init(screen tcell.Screen) error {
	e.fire = New(screen, e.config)
	return e.fire.Init()
}

func (e *FireEffectEffect) Run(quit <-chan struct{}) error {
	return e.fire.Run(quit)
}

func (e *FireEffectEffect) Cleanup() error {
	if e.fire != nil {
		return e.fire.Cleanup()
	}
	return nil
}
