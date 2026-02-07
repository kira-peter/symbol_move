package typewritercode

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// TypewriterCodeEffect 打字机代码雨特效
type TypewriterCodeEffect struct {
	typewriter *Typewriter
	config     *Config
}

// NewEffect 创建打字机代码雨特效实例
func NewEffect() effects.Effect {
	return &TypewriterCodeEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *TypewriterCodeEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "typewriter-code",
		Name:          "打字机代码雨",
		Description:   "模拟打字机逐字显示代码的效果",
		NameEN:        "Typewriter Code",
		DescriptionEN: "Simulates typewriter-style code display character by character",
		LongDescription: `
打字机代码雨特效展示了代码逐字符显示的打字机效果。

特点：
- 多种编程语言代码模板（Go/Python/JS）
- 逐字符打字动画
- 简单语法高亮
- 打字光标闪烁
- 可变打字速度

完美用于：
- 编程主题展示
- 代码演示
- 黑客风格背景
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"代码", "动画", "编程", "打字"},
	}
}

// Init 初始化特效
func (e *TypewriterCodeEffect) Init(screen tcell.Screen) error {
	e.typewriter = New(screen, e.config)
	return e.typewriter.Init()
}

// Run 运行特效
func (e *TypewriterCodeEffect) Run(quit <-chan struct{}) error {
	return e.typewriter.Run(quit)
}

// Cleanup 清理资源
func (e *TypewriterCodeEffect) Cleanup() error {
	if e.typewriter != nil {
		return e.typewriter.Cleanup()
	}
	return nil
}
