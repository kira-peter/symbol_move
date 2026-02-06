package audiovisualizer

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	BarCount int
	FPS      int
}

func DefaultConfig() *Config {
	return &Config{
		BarCount: 40,
		FPS:      30,
	}
}

type AudioVisualizer struct {
	screen   tcell.Screen
	config   *Config
	barHeights []float64
	targetHeights []float64
	width    int
	height   int
	time     float64
	rand     *rand.Rand
	chars    []rune
}

func New(screen tcell.Screen, config *Config) *AudioVisualizer {
	if config == nil {
		config = DefaultConfig()
	}

	return &AudioVisualizer{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		chars:  []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'},
	}
}

func (a *AudioVisualizer) Init() error {
	a.width, a.height = a.screen.Size()
	a.barHeights = make([]float64, a.config.BarCount)
	a.targetHeights = make([]float64, a.config.BarCount)
	a.time = 0

	for i := 0; i < a.config.BarCount; i++ {
		a.barHeights[i] = 0
		a.targetHeights[i] = a.rand.Float64()
	}

	return nil
}

func (a *AudioVisualizer) Update(deltaTime float64) {
	a.time += deltaTime

	// 每隔一段时间更新目标高度
	if int(a.time*10)%5 == 0 {
		for i := 0; i < a.config.BarCount; i++ {
			// 使用正弦波和随机值混合
			sine := math.Sin(float64(i)*0.2 + a.time*2)
			random := a.rand.Float64() * 0.5
			a.targetHeights[i] = (sine + 1) / 2 * 0.7 + random * 0.3
		}
	}

	// 平滑过渡到目标高度
	for i := 0; i < a.config.BarCount; i++ {
		diff := a.targetHeights[i] - a.barHeights[i]
		a.barHeights[i] += diff * deltaTime * 5
	}
}

func (a *AudioVisualizer) Render() {
	a.screen.Clear()

	barWidth := a.width / a.config.BarCount
	if barWidth < 1 {
		barWidth = 1
	}

	for i := 0; i < a.config.BarCount; i++ {
		x := i * barWidth
		barHeight := int(a.barHeights[i] * float64(a.height))

		a.drawBar(x, barWidth, barHeight)
	}

	a.screen.Show()
}

func (a *AudioVisualizer) drawBar(x, width, height int) {
	colors := []tcell.Color{
		tcell.ColorGreen,
		tcell.ColorYellow,
		tcell.ColorOrange,
		tcell.ColorRed,
	}

	for dy := 0; dy < height; dy++ {
		y := a.height - 1 - dy

		// 根据高度选择颜色
		colorIdx := (dy * len(colors)) / a.height
		if colorIdx >= len(colors) {
			colorIdx = len(colors) - 1
		}
		color := colors[colorIdx]

		style := tcell.StyleDefault.Foreground(color)

		for dx := 0; dx < width; dx++ {
			px := x + dx
			if px < a.width && y >= 0 {
				a.screen.SetContent(px, y, '█', nil, style)
			}
		}
	}
}

func (a *AudioVisualizer) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(a.config.FPS))
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

			a.Update(deltaTime)
			a.Render()
		}
	}
}

func (a *AudioVisualizer) Cleanup() error {
	return nil
}
