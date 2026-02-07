package snowfall

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// SnowfallEffect 雪花飘落特效
type SnowfallEffect struct {
	snow   *Snowfall
	config *Config
}

// NewEffect 创建雪花飘落特效实例
func NewEffect() effects.Effect {
	return &SnowfallEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *SnowfallEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "snowfall",
		Name:          "雪花飘落",
		Description:   "浪漫的雪花从天而降",
		NameEN:        "Snowfall",
		DescriptionEN: "Romantic snowflakes falling from above",
		LongDescription: `
雪花飘落特效模拟冬日雪景，展现雪花轻盈飘落的美妙瞬间。

特点：
- 三层深度效果（远、中、近）
- 自然的摆动飘落动画
- 多种雪花形状和大小
- 可调节飘落密度
- 流畅的 30 FPS 动画

完美用于：
- 营造冬日氛围
- 节日主题展示
- 放松心情的背景动画
- 浪漫场景渲染
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"自然", "冬日", "浪漫", "动画"},
	}
}

// Init 初始化特效
func (e *SnowfallEffect) Init(screen tcell.Screen) error {
	e.snow = New(screen, e.config)
	return e.snow.Init()
}

// Run 运行特效
func (e *SnowfallEffect) Run(quit <-chan struct{}) error {
	return e.snow.Run(quit)
}

// Cleanup 清理资源
func (e *SnowfallEffect) Cleanup() error {
	if e.snow != nil {
		return e.snow.Cleanup()
	}
	return nil
}
