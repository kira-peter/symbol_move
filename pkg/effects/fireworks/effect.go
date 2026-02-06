package fireworks

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// FireworksEffect 烟花绽放特效
type FireworksEffect struct {
	fireworks *Fireworks
	config    *Config
}

// NewEffect 创建烟花绽放特效实例
func NewEffect() effects.Effect {
	return &FireworksEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *FireworksEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "fireworks",
		Name:        "烟花绽放",
		Description: "模拟烟花绽放的粒子效果",
		LongDescription: `
烟花绽放特效模拟了真实的烟花效果，包括上升和爆炸两个阶段。

特点：
- 粒子系统和抛物线运动
- 两阶段动画（上升+爆炸）
- 重力模拟和粒子衰减
- 多彩的颜色效果
- 随机发射位置和高度

完美用于：
- 庆祝场景
- 节日氛围
- 粒子效果展示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"粒子", "动画", "庆祝", "多彩"},
	}
}

// Init 初始化特效
func (e *FireworksEffect) Init(screen tcell.Screen) error {
	e.fireworks = New(screen, e.config)
	return e.fireworks.Init()
}

// Run 运行特效
func (e *FireworksEffect) Run(quit <-chan struct{}) error {
	return e.fireworks.Run(quit)
}

// Cleanup 清理资源
func (e *FireworksEffect) Cleanup() error {
	if e.fireworks != nil {
		return e.fireworks.Cleanup()
	}
	return nil
}
