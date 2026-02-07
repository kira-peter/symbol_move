package particleburst

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type ParticleBurstEffect struct {
	burst  *ParticleBurst
	config *Config
}

func NewEffect() effects.Effect {
	return &ParticleBurstEffect{
		config: DefaultConfig(),
	}
}

func (e *ParticleBurstEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "particle-burst",
		Name:          "粒子爆炸",
		Description:   "从中心点向外爆炸的粒子效果,粒子随时间衰减和消失",
		NameEN:        "Particle Burst",
		DescriptionEN: "Particles exploding outward from center with decay over time",
		LongDescription: `
粒子爆炸特效展示壮观的粒子物理动画。

特点：
- 真实的物理模拟
- 重力和衰减效果
- 彩色粒子系统
- 多个爆炸点
- 30 FPS 流畅运行

完美用于：
- 烟花效果
- 粒子系统演示
- 物理模拟展示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"粒子", "爆炸", "动画", "物理"},
	}
}

func (e *ParticleBurstEffect) Init(screen tcell.Screen) error {
	e.burst = New(screen, e.config)
	return e.burst.Init()
}

func (e *ParticleBurstEffect) Run(quit <-chan struct{}) error {
	return e.burst.Run(quit)
}

func (e *ParticleBurstEffect) Cleanup() error {
	if e.burst != nil {
		return e.burst.Cleanup()
	}
	return nil
}
