package bigclock

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type BigClockEffect struct {
	clock  *BigClock
	config *Config
}

func NewEffect() effects.Effect {
	return &BigClockEffect{
		config: DefaultConfig(),
	}
}

func (e *BigClockEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "big-clock",
		Name:        "大字时钟",
		Description: "使用大字符（ASCII Art）显示当前时间，每秒更新",
		LongDescription: `
大字时钟特效使用 ASCII 艺术字显示当前时间。

特点：
- ASCII Art 大字体
- HH:MM:SS 格式
- 每秒自动更新
- 彩色显示
- 居中对齐

完美用于：
- 桌面时钟
- 演示计时器
- 装饰性时钟
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"时钟", "实用", "ASCII艺术", "时间"},
	}
}

func (e *BigClockEffect) Init(screen tcell.Screen) error {
	e.clock = New(screen, e.config)
	return e.clock.Init()
}

func (e *BigClockEffect) Run(quit <-chan struct{}) error {
	return e.clock.Run(quit)
}

func (e *BigClockEffect) Cleanup() error {
	if e.clock != nil {
		return e.clock.Cleanup()
	}
	return nil
}
