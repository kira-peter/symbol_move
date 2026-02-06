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

	// 计算总宽度（每个字符的ASCII艺术宽度）
	totalWidth := 0
	for _, ch := range timeStr {
		if lines, ok := b.digits[ch]; ok && len(lines) > 0 {
			totalWidth += len(lines[0]) // ASCII艺术每个字符占用的宽度
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
				x += len(lines[0])
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
			// 只绘制非空格字符
			if ch != ' ' && py >= 0 && py < b.height && px >= 0 && px < b.width {
				var comb []rune
				b.screen.SetContent(px, py, ch, comb, style)
			}
			// 移动到下一个字符位置（每个ASCII字符占1个宽度）
			px += 1
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
