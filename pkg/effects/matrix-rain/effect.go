package matrixrain

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// MatrixRainEffect 实现 Effect 接口的矩阵字符雨特效
type MatrixRainEffect struct {
	rain   *Rain
	config *Config
}

// NewEffect 创建新的字符雨特效实例（工厂函数）
func NewEffect() effects.Effect {
	return &MatrixRainEffect{
		config: DefaultConfig(),
	}
}

// NewEffectWithConfig 使用自定义配置创建特效实例
func NewEffectWithConfig(config *Config) effects.Effect {
	return &MatrixRainEffect{
		config: config,
	}
}

// Metadata 返回特效元数据
func (e *MatrixRainEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "matrix-rain",
		Name:          "矩阵字符雨",
		Description:   "经典黑客帝国风格字符雨效果",
		NameEN:        "Matrix Rain",
		DescriptionEN: "Classic Matrix-style digital rain effect",
		LongDescription: `经典的"黑客帝国"风格字符雨特效。

特性：
• 多种字符集（数字、字母、日文片假名、混合）
• 颜色渐变效果（亮白→亮绿→绿色→暗绿）
• 可调节速度和密度
• 平滑动画和自适应终端`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"classic", "matrix", "animation", "green"},
	}
}

// Init 初始化特效
func (e *MatrixRainEffect) Init(screen tcell.Screen) error {
	e.rain = New(screen, e.config)
	return nil
}

// Run 运行特效主循环
func (e *MatrixRainEffect) Run(quit <-chan struct{}) error {
	if e.rain == nil {
		return nil
	}

	ticker := time.NewTicker(time.Second / time.Duration(e.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			e.rain.Update()
			e.rain.Render()
		}
	}
}

// Cleanup 清理资源
func (e *MatrixRainEffect) Cleanup() error {
	// 字符雨效果无需特殊清理
	e.rain = nil
	return nil
}

// SetConfig 设置配置（便捷方法）
func (e *MatrixRainEffect) SetConfig(config *Config) {
	e.config = config
}

// GetConfig 获取配置（便捷方法）
func (e *MatrixRainEffect) GetConfig() *Config {
	return e.config
}
