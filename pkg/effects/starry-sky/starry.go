package starrysky

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Density 星星密度
type Density int

const (
	DensitySparse Density = iota // 稀疏 ~1%
	DensityMedium                 // 中等 ~2%
	DensityDense                  // 密集 ~3%
)

// Theme 颜色主题
type Theme int

const (
	ThemeClassic Theme = iota // 经典白色
	ThemeColorful             // 彩色
	ThemeBlue                 // 蓝色主题
)

// Config 星空配置
type Config struct {
	Density Density // 星星密度
	Theme   Theme   // 颜色主题
	FPS     int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Density: DensityMedium,
		Theme:   ThemeClassic,
		FPS:     30,
	}
}

// Star 单个星星
type Star struct {
	x, y      int       // 位置
	char      rune      // 字符
	baseColor tcell.Color // 基础颜色
	brightness float64  // 当前亮度 (0.0-1.0)
	phase     float64   // 闪烁相位 (0-2π)
	speed     float64   // 闪烁速度
}

// StarrySky 星空特效
type StarrySky struct {
	screen     tcell.Screen
	config     *Config
	stars      []*Star
	width      int
	height     int
	rand       *rand.Rand
	lastUpdate time.Time
}

// New 创建星空特效实例
func New(screen tcell.Screen, config *Config) *StarrySky {
	if config == nil {
		config = DefaultConfig()
	}

	return &StarrySky{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化星空
func (s *StarrySky) Init() error {
	s.width, s.height = s.screen.Size()
	s.generateStars()
	s.lastUpdate = time.Now()
	return nil
}

// generateStars 生成星星
func (s *StarrySky) generateStars() {
	totalCells := s.width * s.height

	// 根据密度计算星星数量
	var coverage float64
	switch s.config.Density {
	case DensitySparse:
		coverage = 0.01
	case DensityMedium:
		coverage = 0.02
	case DensityDense:
		coverage = 0.03
	}

	starCount := int(float64(totalCells) * coverage)
	s.stars = make([]*Star, 0, starCount)

	// 生成星星
	for i := 0; i < starCount; i++ {
		star := &Star{
			x:          s.rand.Intn(s.width),
			y:          s.rand.Intn(s.height),
			char:       s.randomStarChar(),
			baseColor:  s.randomStarColor(),
			brightness: s.rand.Float64(), // 随机初始亮度
			phase:      s.rand.Float64() * 2 * math.Pi, // 随机初始相位
			speed:      0.5 + s.rand.Float64()*1.5, // 0.5-2.0 速度倍数
		}
		s.stars = append(s.stars, star)
	}
}

// randomStarChar 随机选择星星字符
func (s *StarrySky) randomStarChar() rune {
	chars := []rune{'*', '·', '.', '+', '✦', '✧'}
	return chars[s.rand.Intn(len(chars))]
}

// randomStarColor 根据主题随机选择星星颜色
func (s *StarrySky) randomStarColor() tcell.Color {
	switch s.config.Theme {
	case ThemeClassic:
		return tcell.ColorWhite
	case ThemeColorful:
		colors := []tcell.Color{
			tcell.ColorLightBlue,
			tcell.ColorLightYellow,
			tcell.ColorLightCyan,
			tcell.ColorWhite,
		}
		return colors[s.rand.Intn(len(colors))]
	case ThemeBlue:
		colors := []tcell.Color{
			tcell.ColorBlue,
			tcell.ColorLightBlue,
			tcell.ColorDarkBlue,
			tcell.ColorCadetBlue,
		}
		return colors[s.rand.Intn(len(colors))]
	}
	return tcell.ColorWhite
}

// Update 更新星空状态
func (s *StarrySky) Update(deltaTime float64) {
	for _, star := range s.stars {
		// 更新闪烁相位
		star.phase += deltaTime * star.speed
		if star.phase > 2*math.Pi {
			star.phase -= 2 * math.Pi
		}

		// 使用正弦波计算亮度 (0.3-1.0 范围，避免完全暗)
		star.brightness = 0.3 + 0.7*(0.5+0.5*math.Sin(star.phase))
	}
}

// Render 渲染星空
func (s *StarrySky) Render() {
	s.screen.Clear()

	for _, star := range s.stars {
		style := s.getStarStyle(star)
		s.screen.SetContent(star.x, star.y, star.char, nil, style)
	}

	s.screen.Show()
}

// getStarStyle 根据亮度获取星星样式
func (s *StarrySky) getStarStyle(star *Star) tcell.Style {
	// 根据亮度选择颜色深浅
	var color tcell.Color

	if star.brightness > 0.8 {
		// 非常亮 - 使用亮白或基础颜色的亮版本
		if star.baseColor == tcell.ColorWhite {
			color = tcell.ColorWhite
		} else {
			color = star.baseColor
		}
	} else if star.brightness > 0.5 {
		// 中等亮度 - 使用基础颜色
		color = star.baseColor
	} else {
		// 暗淡 - 使用暗色版本
		color = tcell.ColorGray
	}

	style := tcell.StyleDefault.Foreground(color)

	// 最亮的星星加粗
	if star.brightness > 0.9 {
		style = style.Bold(true)
	}

	return style
}

// Run 运行星空特效
func (s *StarrySky) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(s.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(s.lastUpdate).Seconds()
			s.lastUpdate = now

			s.Update(deltaTime)
			s.Render()
		}
	}
}

// Cleanup 清理资源
func (s *StarrySky) Cleanup() error {
	// 清理资源（如果需要）
	return nil
}
