package dnahelix

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config DNA双螺旋配置
type Config struct {
	RotationSpeed float64 // 旋转速度 (弧度/秒)
	HelixRadius   float64 // 螺旋半径
	BaseSpacing   float64 // 碱基对间距
	FPS           int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		RotationSpeed: 1.0,  // 1弧度/秒
		HelixRadius:   8.0,  // 8个字符宽度
		BaseSpacing:   0.15, // 螺旋紧密度
		FPS:           30,
	}
}

// DNAHelix DNA双螺旋特效
type DNAHelix struct {
	screen     tcell.Screen
	config     *Config
	angle      float64
	width      int
	height     int
	lastUpdate time.Time
}

// New 创建DNA双螺旋特效实例
func New(screen tcell.Screen, config *Config) *DNAHelix {
	if config == nil {
		config = DefaultConfig()
	}

	return &DNAHelix{
		screen: screen,
		config: config,
	}
}

// Init 初始化DNA双螺旋
func (d *DNAHelix) Init() error {
	d.width, d.height = d.screen.Size()
	d.angle = 0
	d.lastUpdate = time.Now()
	return nil
}

// Update 更新DNA双螺旋状态
func (d *DNAHelix) Update(deltaTime float64) {
	d.angle += d.config.RotationSpeed * deltaTime
	// 保持角度在合理范围内
	if d.angle > 2*math.Pi {
		d.angle -= 2 * math.Pi
	}
}

// Render 渲染DNA双螺旋
func (d *DNAHelix) Render() {
	d.screen.Clear()

	centerX := d.width / 2
	centerY := d.height / 2

	// 颜色定义
	styleBlue := tcell.StyleDefault.Foreground(tcell.ColorLightBlue).Bold(true)
	styleRed := tcell.StyleDefault.Foreground(tcell.ColorLightCoral).Bold(true)
	styleGray := tcell.StyleDefault.Foreground(tcell.ColorGray)
	styleYellow := tcell.StyleDefault.Foreground(tcell.ColorYellow)

	// 绘制螺旋
	for z := -d.height / 2; z < d.height/2; z++ {
		y := centerY + z

		if y < 0 || y >= d.height {
			continue
		}

		// 计算两条螺旋的位置
		theta := d.angle + float64(z)*d.config.BaseSpacing

		x1 := centerX + int(d.config.HelixRadius*math.Cos(theta))
		x2 := centerX + int(d.config.HelixRadius*math.Cos(theta+math.Pi))

		// 确保坐标在屏幕范围内
		if x1 >= 0 && x1 < d.width {
			d.screen.SetContent(x1, y, '●', nil, styleBlue)
		}
		if x2 >= 0 && x2 < d.width {
			d.screen.SetContent(x2, y, '●', nil, styleRed)
		}

		// 绘制碱基对连接线（每3行一次）
		if z%3 == 0 {
			minX := x1
			maxX := x2
			if x1 > x2 {
				minX = x2
				maxX = x1
			}

			// 计算连接线的深度（用于颜色）
			sinVal := math.Sin(theta + math.Pi/2)

			for x := minX; x <= maxX; x++ {
				if x >= 0 && x < d.width {
					// 中点附近画碱基对
					if x == (minX+maxX)/2 {
						// 根据深度选择碱基对颜色
						if sinVal > 0 {
							d.screen.SetContent(x, y, '═', nil, styleYellow)
						} else {
							d.screen.SetContent(x, y, '═', nil, styleGray)
						}
					} else {
						// 其他部分画连接线
						d.screen.SetContent(x, y, '─', nil, styleGray)
					}
				}
			}
		}
	}

	d.screen.Show()
}

// Run 运行DNA双螺旋特效
func (d *DNAHelix) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(d.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(d.lastUpdate).Seconds()
			d.lastUpdate = now

			d.Update(deltaTime)
			d.Render()
		}
	}
}

// Cleanup 清理资源
func (d *DNAHelix) Cleanup() error {
	return nil
}
