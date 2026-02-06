package heartbeat

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 心跳配置
type Config struct {
	BPM      int     // 心跳次数/分钟
	MaxScale float64 // 最大缩放倍数
	FPS      int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		BPM:      72,   // 正常心率
		MaxScale: 0.3,  // 30%的缩放范围
		FPS:      30,
	}
}

// Heartbeat 心跳特效
type Heartbeat struct {
	screen     tcell.Screen
	config     *Config
	time       float64
	width      int
	height     int
	asciiArt   []string
	lastUpdate time.Time
}

// heartASCII 心形ASCII艺术
var heartASCII = []string{
	"    ██   ██    ",
	"  ██ ██ ██ ██  ",
	" ██   ███   ██ ",
	"██           ██",
	" ██         ██ ",
	"  ██       ██  ",
	"   ██     ██   ",
	"    ██   ██    ",
	"     ██ ██     ",
	"      ███      ",
	"       █       ",
}

// New 创建心跳特效实例
func New(screen tcell.Screen, config *Config) *Heartbeat {
	if config == nil {
		config = DefaultConfig()
	}

	return &Heartbeat{
		screen:   screen,
		config:   config,
		asciiArt: heartASCII,
	}
}

// Init 初始化心跳特效
func (h *Heartbeat) Init() error {
	h.width, h.height = h.screen.Size()
	h.time = 0
	h.lastUpdate = time.Now()
	return nil
}

// Update 更新心跳状态
func (h *Heartbeat) Update(deltaTime float64) {
	h.time += deltaTime
}

// Render 渲染心跳
func (h *Heartbeat) Render() {
	h.screen.Clear()

	// 计算当前缩放
	beatPhase := 2 * math.Pi * float64(h.config.BPM) / 60.0 * h.time
	scale := 1.0 + h.config.MaxScale*math.Sin(beatPhase)

	// 计算居中位置
	centerX := h.width / 2
	centerY := h.height / 2
	artHeight := len(h.asciiArt)
	artWidth := len(h.asciiArt[0])

	// 绘制缩放后的心形
	for i, line := range h.asciiArt {
		scaledY := centerY + int((float64(i)-float64(artHeight)/2)*scale)

		if scaledY < 0 || scaledY >= h.height {
			continue
		}

		for j, ch := range line {
			if ch == ' ' {
				continue
			}

			scaledX := centerX + int((float64(j)-float64(artWidth)/2)*scale)

			if scaledX < 0 || scaledX >= h.width {
				continue
			}

			// 颜色根据缩放变化
			var color tcell.Color
			if scale > 1.15 {
				color = tcell.ColorRed
			} else if scale > 1.05 {
				color = tcell.ColorLightCoral
			} else {
				color = tcell.ColorDarkRed
			}

			style := tcell.StyleDefault.Foreground(color).Bold(true)
			h.screen.SetContent(scaledX, scaledY, ch, nil, style)
		}
	}

	h.screen.Show()
}

// Run 运行心跳特效
func (h *Heartbeat) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(h.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(h.lastUpdate).Seconds()
			h.lastUpdate = now

			h.Update(deltaTime)
			h.Render()
		}
	}
}

// Cleanup 清理资源
func (h *Heartbeat) Cleanup() error {
	return nil
}
