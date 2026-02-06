package wavetext

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 波浪文字配置
type Config struct {
	Text       string  // 显示文本
	Amplitude  float64 // 波浪振幅（字符高度）
	WaveSpeed  float64 // 波浪速度
	ColorSpeed float64 // 颜色变化速度
	FPS        int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Text:       "SymbolMove",
		Amplitude:  3.0,
		WaveSpeed:  2.0,
		ColorSpeed: 1.0,
		FPS:        30,
	}
}

// WaveText 波浪文字特效
type WaveText struct {
	screen     tcell.Screen
	config     *Config
	phase      float64 // 波浪相位
	colorPhase float64 // 颜色相位
	width      int
	height     int
}

// New 创建波浪文字特效实例
func New(screen tcell.Screen, config *Config) *WaveText {
	if config == nil {
		config = DefaultConfig()
	}

	return &WaveText{
		screen: screen,
		config: config,
	}
}

// Init 初始化波浪文字
func (w *WaveText) Init() error {
	w.width, w.height = w.screen.Size()
	w.phase = 0
	w.colorPhase = 0
	return nil
}

// Update 更新波浪文字状态
func (w *WaveText) Update(deltaTime float64) {
	// 更新波浪相位
	w.phase += deltaTime * w.config.WaveSpeed
	if w.phase > 2*math.Pi {
		w.phase -= 2 * math.Pi
	}

	// 更新颜色相位
	w.colorPhase += deltaTime * w.config.ColorSpeed
	if w.colorPhase > 360 {
		w.colorPhase -= 360
	}
}

// Render 渲染波浪文字
func (w *WaveText) Render() {
	w.screen.Clear()

	text := []rune(w.config.Text)
	textLen := len(text)

	// 计算起始 X 坐标（居中）
	startX := (w.width - textLen) / 2
	if startX < 0 {
		startX = 0
	}

	// 中心 Y 坐标
	centerY := w.height / 2

	// 相位差（相邻字符之间）
	phaseShift := math.Pi / 4

	// 颜色差（相邻字符之间）
	colorShift := 360.0 / float64(textLen)

	// 渲染每个字符
	for i, char := range text {
		x := startX + i
		if x >= w.width {
			break
		}

		// 计算 Y 坐标（正弦波）
		offset := w.config.Amplitude * math.Sin(w.phase+float64(i)*phaseShift)
		y := centerY + int(offset)

		// 确保 Y 坐标在屏幕内
		if y >= 0 && y < w.height {
			// 计算颜色
			hue := math.Mod(w.colorPhase+float64(i)*colorShift, 360)
			color := w.hueToColor(hue)
			style := tcell.StyleDefault.Foreground(color).Bold(true)

			w.screen.SetContent(x, y, char, nil, style)
		}
	}

	w.screen.Show()
}

// hueToColor 将色相转换为 tcell 颜色
func (w *WaveText) hueToColor(hue float64) tcell.Color {
	// 简化的彩虹颜色映射
	h := int(hue)
	switch {
	case h < 60:
		return tcell.ColorRed
	case h < 120:
		return tcell.ColorYellow
	case h < 180:
		return tcell.ColorGreen
	case h < 240:
		return tcell.ColorLightBlue
	case h < 300:
		return tcell.ColorBlue
	default:
		return tcell.ColorPurple
	}
}

// Run 运行波浪文字特效
func (w *WaveText) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(w.config.FPS))
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

			w.Update(deltaTime)
			w.Render()
		}
	}
}

// Cleanup 清理资源
func (w *WaveText) Cleanup() error {
	return nil
}
