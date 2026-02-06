package fireeffect

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	Intensity float64
	FPS       int
}

func DefaultConfig() *Config {
	return &Config{
		Intensity: 1.0,
		FPS:       30,
	}
}

type FireEffect struct {
	screen  tcell.Screen
	config  *Config
	buffer  [][]float64
	width   int
	height  int
	rand    *rand.Rand
	chars   []rune
}

func New(screen tcell.Screen, config *Config) *FireEffect {
	if config == nil {
		config = DefaultConfig()
	}

	return &FireEffect{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		chars:  []rune{' ', '.', ':', '*', 's', 'S', '#', '$', '@'},
	}
}

func (f *FireEffect) Init() error {
	f.width, f.height = f.screen.Size()
	f.buffer = make([][]float64, f.height)
	for i := range f.buffer {
		f.buffer[i] = make([]float64, f.width)
	}
	return nil
}

func (f *FireEffect) Update() {
	// 底部热源
	for x := 0; x < f.width; x++ {
		f.buffer[f.height-1][x] = f.rand.Float64() * f.config.Intensity
	}

	// 火焰向上传播和冷却
	for y := 0; y < f.height-1; y++ {
		for x := 0; x < f.width; x++ {
			// 收集下方和周围的热量
			heat := 0.0
			count := 0

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					ny, nx := y+dy+1, x+dx
					if ny >= 0 && ny < f.height && nx >= 0 && nx < f.width {
						heat += f.buffer[ny][nx]
						count++
					}
				}
			}

			// 平均热量并冷却
			if count > 0 {
				f.buffer[y][x] = (heat / float64(count)) * 0.95
			}
		}
	}
}

func (f *FireEffect) Render() {
	f.screen.Clear()

	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			heat := f.buffer[y][x]
			if heat > 0.05 {
				char, color := f.heatToChar(heat)
				style := tcell.StyleDefault.Foreground(color)
				f.screen.SetContent(x, y, char, nil, style)
			}
		}
	}

	f.screen.Show()
}

func (f *FireEffect) heatToChar(heat float64) (rune, tcell.Color) {
	idx := int(heat * float64(len(f.chars)))
	if idx >= len(f.chars) {
		idx = len(f.chars) - 1
	}

	char := f.chars[idx]

	var color tcell.Color
	if heat > 0.8 {
		color = tcell.ColorYellow
	} else if heat > 0.5 {
		color = tcell.ColorOrange
	} else if heat > 0.3 {
		color = tcell.ColorRed
	} else {
		color = tcell.ColorDarkRed
	}

	return char, color
}

func (f *FireEffect) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(f.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			f.Update()
			f.Render()
		}
	}
}

func (f *FireEffect) Cleanup() error {
	return nil
}
