package waterripple

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// WaterRippleEffect 水波涟漪特效
type WaterRippleEffect struct {
	ripple *WaterRipple
	config *Config
}

// NewEffect 创建水波涟漪特效实例
func NewEffect() effects.Effect {
	return &WaterRippleEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *WaterRippleEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "water-ripple",
		Name:          "水波涟漪",
		Description:   "模拟水滴落下形成的涟漪效果",
		NameEN:        "Water Ripple",
		DescriptionEN: "Simulates ripple effects from water drops",
		LongDescription: `
水波涟漪特效模拟了水滴落入水面形成的波纹扩散效果。

特点：
- 基于波动方程的物理模拟
- 多个水滴波动叠加
- 波动衰减效果
- 干涉现象可视化
- 随机水滴位置

完美用于：
- 放松心情
- 物理演示
- 自然效果展示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"物理", "动画", "自然", "放松"},
	}
}

// Init 初始化特效
func (e *WaterRippleEffect) Init(screen tcell.Screen) error {
	e.ripple = New(screen, e.config)
	return e.ripple.Init()
}

// Run 运行特效
func (e *WaterRippleEffect) Run(quit <-chan struct{}) error {
	return e.ripple.Run(quit)
}

// Cleanup 清理资源
func (e *WaterRippleEffect) Cleanup() error {
	if e.ripple != nil {
		return e.ripple.Cleanup()
	}
	return nil
}
