package bigclock

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	Color tcell.Color
	FPS   int
}

func DefaultConfig() *Config {
	return &Config{
		Color: tcell.ColorLightBlue,
		FPS:   1,
	}
}

type BigClock struct {
	screen tcell.Screen
	config *Config
	width  int
	height int
	digits map[rune][]string
}

func New(screen tcell.Screen, config *Config) *BigClock {
	if config == nil {
		config = DefaultConfig()
	}

	return &BigClock{
		screen: screen,
		config: config,
		digits: initDigits(),
	}
}

func initDigits() map[rune][]string {
	return map[rune][]string{
		'0': {
			" ███ ",
			"█   █",
			"█   █",
			"█   █",
			" ███ ",
		},
		'1': {
			"  █  ",
			" ██  ",
			"  █  ",
			"  █  ",
			" ███ ",
		},
		'2': {
			" ███ ",
			"    █",
			" ███ ",
			"█    ",
			"█████",
		},
		'3': {
			" ███ ",
			"    █",
			"  ██ ",
			"    █",
			" ███ ",
		},
		'4': {
			"█   █",
			"█   █",
			"█████",
			"    █",
			"    █",
		},
		'5': {
			"█████",
			"█    ",
			"████ ",
			"    █",
			"████ ",
		},
		'6': {
			" ███ ",
			"█    ",
			"████ ",
			"█   █",
			" ███ ",
		},
		'7': {
			"█████",
			"    █",
			"   █ ",
			"  █  ",
			"  █  ",
		},
		'8': {
			" ███ ",
			"█   █",
			" ███ ",
			"█   █",
			" ███ ",
		},
		'9': {
			" ███ ",
			"█   █",
			" ████",
			"    █",
			" ███ ",
		},
		':': {
			"     ",
			"  █  ",
			"     ",
			"  █  ",
			"     ",
		},
	}
}

func (b *BigClock) Init() error {
	b.width, b.height = b.screen.Size()
	return nil
}

func (b *BigClock) Render() {
	b.screen.Clear()

	// 获取当前时间
	now := time.Now()
	timeStr := now.Format("15:04:05")

	// 计算每个字符的实际显示宽度
	calcWidth := func(s string) int {
		w := 0
		for _, ch := range s {
			if ch > 127 {
				w += 2
			} else {
				w += 1
			}
		}
		return w
	}

	// 计算总宽度
	totalWidth := 0
	for _, ch := range timeStr {
		if lines, ok := b.digits[ch]; ok && len(lines) > 0 {
			totalWidth += calcWidth(lines[0])
		}
	}

	startX := (b.width - totalWidth) / 2
	startY := (b.height - 5) / 2

	// 渲染每个字符
	x := startX
	for _, ch := range timeStr {
		if lines, ok := b.digits[ch]; ok {
			b.renderDigit(x, startY, lines)
			// 移动到下一个字符位置
			if len(lines) > 0 {
				x += calcWidth(lines[0])
			}
		}
	}

	b.screen.Show()
}

func (b *BigClock) renderDigit(x, y int, lines []string) {
	style := tcell.StyleDefault.Foreground(b.config.Color).Bold(true)

	for row, line := range lines {
		px := x
		for _, ch := range line {
			py := y + row
			if py >= 0 && py < b.height && px >= 0 && px < b.width {
				// 绘制所有字符（包括空格），保持布局正确
				var comb []rune
				b.screen.SetContent(px, py, ch, comb, style)
			}
			// 中文字符占2个宽度
			if ch > 127 {
				px += 2
			} else {
				px += 1
			}
		}
	}
}

func (b *BigClock) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(b.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			b.Render()
		}
	}
}

func (b *BigClock) Cleanup() error {
	return nil
}
