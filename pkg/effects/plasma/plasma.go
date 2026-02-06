package plasma

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	Speed float64
	FPS   int
}

func DefaultConfig() *Config {
	return &Config{
		Speed: 1.0,
		FPS:   30,
	}
}

type Plasma struct {
	screen tcell.Screen
	config *Config
	width  int
	height int
	time   float64
	chars  []rune
	colors []tcell.Color
}

func New(screen tcell.Screen, config *Config) *Plasma {
	if config == nil {
		config = DefaultConfig()
	}

	return &Plasma{
		screen: screen,
		config: config,
		chars:  []rune{' ', '░', '▒', '▓', '█'},
		colors: []tcell.Color{
			tcell.ColorBlue,
			tcell.ColorLightBlue,
			tcell.ColorGreen,
			tcell.ColorYellow,
			tcell.ColorRed,
			tcell.ColorPink,
			tcell.ColorPurple,
		},
	}
}

func (p *Plasma) Init() error {
	p.width, p.height = p.screen.Size()
	p.time = 0
	return nil
}

func (p *Plasma) Update(deltaTime float64) {
	p.time += deltaTime * p.config.Speed
}

func (p *Plasma) plasmaValue(x, y int) float64 {
	fx := float64(x) / float64(p.width)
	fy := float64(y) / float64(p.height)

	// 多层正弦波叠加
	v := 0.0
	v += math.Sin((fx*10 + p.time))
	v += math.Sin((fy*10 + p.time))
	v += math.Sin((fx*10 + fy*10 + p.time) / 2)

	cx := fx + 0.5*math.Sin(p.time/5)
	cy := fy + 0.5*math.Cos(p.time/3)
	v += math.Sin(math.Sqrt(100*(cx*cx+cy*cy)) + p.time)

	// 归一化到 0-1
	return (v + 4) / 8
}

func (p *Plasma) Render() {
	p.screen.Clear()

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			value := p.plasmaValue(x, y)

			// 选择字符
			charIdx := int(value * float64(len(p.chars)))
			if charIdx >= len(p.chars) {
				charIdx = len(p.chars) - 1
			}
			char := p.chars[charIdx]

			// 选择颜色
			colorIdx := int(value * float64(len(p.colors)))
			if colorIdx >= len(p.colors) {
				colorIdx = len(p.colors) - 1
			}
			color := p.colors[colorIdx]

			style := tcell.StyleDefault.Foreground(color)
			p.screen.SetContent(x, y, char, nil, style)
		}
	}

	p.screen.Show()
}

func (p *Plasma) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(p.config.FPS))
	defer ticker.Stop()

	lastUpdate := time.Now()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(lastUpdate).Seconds()
			lastUpdate = now

			p.Update(deltaTime)
			p.Render()
		}
	}
}

func (p *Plasma) Cleanup() error {
	return nil
}
