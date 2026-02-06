package matrixtunnel

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 矩阵隧道配置
type Config struct {
	Speed   float64 // 飞行速度
	Density float64 // 字符密度
	FPS     int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Speed:   5.0,  // 飞行速度
		Density: 0.3,  // 字符密度
		FPS:     30,
	}
}

// MatrixTunnel 矩阵隧道特效
type MatrixTunnel struct {
	screen     tcell.Screen
	config     *Config
	depth      float64
	width      int
	height     int
	chars      []rune
	lastUpdate time.Time
	rand       *rand.Rand
}

// New 创建矩阵隧道特效实例
func New(screen tcell.Screen, config *Config) *MatrixTunnel {
	if config == nil {
		config = DefaultConfig()
	}

	// 矩阵字符集
	chars := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'Z', 'T', 'M', 'X',
		'ﾊ', 'ﾐ', 'ﾋ', 'ｰ', 'ｳ', 'ｼ', 'ﾅ', 'ﾓ', 'ﾆ', 'ｻ',
	}

	return &MatrixTunnel{
		screen: screen,
		config: config,
		chars:  chars,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化矩阵隧道
func (m *MatrixTunnel) Init() error {
	m.width, m.height = m.screen.Size()
	m.depth = 0
	m.lastUpdate = time.Now()
	return nil
}

// Update 更新矩阵隧道状态
func (m *MatrixTunnel) Update(deltaTime float64) {
	m.depth += m.config.Speed * deltaTime
	// 循环深度
	if m.depth > 50.0 {
		m.depth -= 50.0
	}
}

// getColorByBrightness 根据亮度获取颜色
func (m *MatrixTunnel) getColorByBrightness(brightness float64) tcell.Color {
	if brightness > 0.7 {
		return tcell.ColorGreen
	} else if brightness > 0.4 {
		return tcell.ColorLightGreen
	} else if brightness > 0.2 {
		return tcell.ColorDarkGreen
	}
	return tcell.ColorBlack
}

// Render 渲染矩阵隧道
func (m *MatrixTunnel) Render() {
	m.screen.Clear()

	centerX := m.width / 2
	centerY := m.height / 2
	tunnelRadius := 15.0 // 隧道半径

	// 绘制隧道墙壁
	for z := 1.0; z < 50.0; z += 0.5 {
		actualDepth := z + m.depth
		// 循环深度
		if actualDepth > 50.0 {
			actualDepth -= 50.0
		}

		// 根据密度决定是否绘制
		if m.rand.Float64() > m.config.Density {
			continue
		}

		// 圆形断面
		angle := m.rand.Float64() * 2 * math.Pi
		radius := tunnelRadius * (0.8 + m.rand.Float64()*0.4)

		x := radius * math.Cos(angle)
		y := radius * math.Sin(angle)

		// 透视投影
		// 距离越近（z越小），投影越大
		scale := 30.0 / actualDepth
		screenX := centerX + int(x*scale)
		screenY := centerY + int(y*scale)

		// 检查边界
		if screenX >= 0 && screenX < m.width && screenY >= 0 && screenY < m.height {
			// 根据深度选择字符和颜色
			brightness := 1.0 / actualDepth
			char := m.chars[int(actualDepth*10+angle*5)%len(m.chars)]

			color := m.getColorByBrightness(brightness)
			style := tcell.StyleDefault.Foreground(color)

			// 最亮的字符加粗
			if brightness > 0.8 {
				style = style.Bold(true)
			}

			m.screen.SetContent(screenX, screenY, char, nil, style)
		}
	}

	m.screen.Show()
}

// Run 运行矩阵隧道特效
func (m *MatrixTunnel) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(m.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(m.lastUpdate).Seconds()
			m.lastUpdate = now

			m.Update(deltaTime)
			m.Render()
		}
	}
}

// Cleanup 清理资源
func (m *MatrixTunnel) Cleanup() error {
	return nil
}
