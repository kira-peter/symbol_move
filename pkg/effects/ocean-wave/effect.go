package oceanwave

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// OceanWaveEffect 字符海浪特效
type OceanWaveEffect struct {
	ocean  *OceanWave
	config *Config
}

// NewEffect 创建字符海浪特效实例
func NewEffect() effects.Effect {
	return &OceanWaveEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *OceanWaveEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "ocean-wave",
		Name:          "字符海浪",
		Description:   "模拟海浪波动的自然效果",
		NameEN:        "Ocean Wave",
		DescriptionEN: "Simulates natural ocean wave motion",
		LongDescription: `
字符海浪特效展示了海洋的波涛起伏，多层波浪叠加形成真实的海浪效果。

特点：
- 多层波浪叠加（不同频率和振幅）
- 浪花效果和泡沫字符
- 深蓝到白色的渐变色
- 流畅的波浪动画
- 深度感的颜色过渡

完美用于：
- 放松心情的背景
- 自然主题展示
- 海洋相关应用
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"自然", "动画", "放松", "海洋"},
	}
}

// Init 初始化特效
func (e *OceanWaveEffect) Init(screen tcell.Screen) error {
	e.ocean = New(screen, e.config)
	return e.ocean.Init()
}

// Run 运行特效
func (e *OceanWaveEffect) Run(quit <-chan struct{}) error {
	return e.ocean.Run(quit)
}

// Cleanup 清理资源
func (e *OceanWaveEffect) Cleanup() error {
	if e.ocean != nil {
		return e.ocean.Cleanup()
	}
	return nil
}
