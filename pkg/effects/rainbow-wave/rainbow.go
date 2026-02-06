package rainbowwave

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 彩虹波浪配置
type Config struct {
	WaveSpeed  float64 // 波浪速度
	WaveHeight float64 // 波浪高度
	NumWaves   int     // 波浪数量
	FPS        int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		WaveSpeed:  1.0,
		WaveHeight: 5.0,
		NumWaves:   3,
		FPS:        30,
	}
}

// RainbowWave 彩虹波浪特效
type RainbowWave struct {
	screen tcell.Screen
	config *Config
	phase  float64
	width  int
	height int
	colors []tcell.Color
	chars  []rune
}

// New 创建彩虹波浪特效实例
func New(screen tcell.Screen, config *Config) *RainbowWave {
	if config == nil {
		config = DefaultConfig()
	}

	return &RainbowWave{
		screen: screen,
		config: config,
		colors: []tcell.Color{
			tcell.ColorRed,
			tcell.ColorOrange,
			tcell.ColorYellow,
			tcell.ColorGreen,
			tcell.ColorLightBlue,
			tcell.ColorBlue,
			tcell.ColorPurple,
		},
		chars: []rune{'~', '≈', '∼', '≋'},
	}
}

// Init 初始化彩虹波浪
func (r *RainbowWave) Init() error {
	r.width, r.height = r.screen.Size()
	r.phase = 0
	return nil
}

// Update 更新彩虹波浪状态
func (r *RainbowWave) Update(deltaTime float64) {
	r.phase += deltaTime * r.config.WaveSpeed
	if r.phase > 2*math.Pi {
		r.phase -= 2 * math.Pi
	}
}

// Render 渲染彩虹波浪
func (r *RainbowWave) Render() {
	r.screen.Clear()

	centerY := r.height / 2

	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			// 计算波浪值
			waveValue := 0.0
			for i := 0; i < r.config.NumWaves; i++ {
				freq := float64(i+1) * 0.5
				amp := r.config.WaveHeight / float64(i+1)
				waveValue += amp * math.Sin(r.phase+float64(x)*0.1*freq)
			}

			waveY := centerY + int(waveValue)

			// 如果当前位置在波浪范围内
			if y >= waveY-1 && y <= waveY+1 {
				// 根据 X 坐标选择颜色
				colorIdx := (x + int(r.phase*10)) % len(r.colors)
				color := r.colors[colorIdx]

				// 根据距离中心的距离选择字符
				charIdx := abs(y-waveY) % len(r.chars)
				char := r.chars[charIdx]

				style := tcell.StyleDefault.Foreground(color)
				r.screen.SetContent(x, y, char, nil, style)
			}
		}
	}

	r.screen.Show()
}

// abs 返回整数的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Run 运行彩虹波浪特效
func (r *RainbowWave) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(r.config.FPS))
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

			r.Update(deltaTime)
			r.Render()
		}
	}
}

// Cleanup 清理资源
func (r *RainbowWave) Cleanup() error {
	return nil
}
