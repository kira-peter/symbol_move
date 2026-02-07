package digitalwaterfall

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// DigitalWaterfallEffect 数字瀑布特效
type DigitalWaterfallEffect struct {
	waterfall *DigitalWaterfall
	config    *Config
}

// NewEffect 创建数字瀑布特效实例
func NewEffect() effects.Effect {
	return &DigitalWaterfallEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *DigitalWaterfallEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "digital-waterfall",
		Name:          "数字瀑布",
		Description:   "数字0-9从上到下快速流动,绿色主题的数字瀑布效果",
		NameEN:        "Digital Waterfall",
		DescriptionEN: "Numbers 0-9 flowing downward in a green-themed cascade",
		LongDescription: `
数字瀑布特效模拟黑客帝国风格的数字流动效果。

特点：
- 只使用数字字符 0-9
- 绿色科技主题
- 流畅的流动动画
- 渐变尾迹效果
- 随机白色闪光
- 30 FPS 流畅运行
- 自动适配终端大小

完美用于：
- 黑客主题展示
- 科技感背景动画
- 终端演示效果
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"科技", "动画", "绿色", "经典"},
	}
}

// Init 初始化特效
func (e *DigitalWaterfallEffect) Init(screen tcell.Screen) error {
	e.waterfall = New(screen, e.config)
	return e.waterfall.Init()
}

// Run 运行特效
func (e *DigitalWaterfallEffect) Run(quit <-chan struct{}) error {
	return e.waterfall.Run(quit)
}

// Cleanup 清理资源
func (e *DigitalWaterfallEffect) Cleanup() error {
	if e.waterfall != nil {
		return e.waterfall.Cleanup()
	}
	return nil
}
