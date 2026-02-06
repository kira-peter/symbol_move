package snowfall

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Density 雪花密度
type Density int

const (
	DensitySparse Density = iota // 稀疏
	DensityMedium                 // 中等
	DensityDense                  // 密集
)

// Config 雪花配置
type Config struct {
	Density Density // 雪花密度
	FPS     int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Density: DensityMedium,
		FPS:     30,
	}
}

// Snowflake 单个雪花
type Snowflake struct {
	x, y       float64     // 位置（浮点数用于平滑移动）
	char       rune        // 字符
	color      tcell.Color // 颜色
	speed      float64     // 下落速度
	swingPhase float64     // 摆动相位
	swingAmp   float64     // 摆动幅度
	layer      int         // 层次 (0=远, 1=中, 2=近)
}

// Snowfall 雪花飘落特效
type Snowfall struct {
	screen     tcell.Screen
	config     *Config
	flakes     []*Snowflake
	width      int
	height     int
	rand       *rand.Rand
	lastUpdate time.Time
}

const (
	maxFlakes = 500 // 最大雪花数量
)

// New 创建雪花飘落实例
func New(screen tcell.Screen, config *Config) *Snowfall {
	if config == nil {
		config = DefaultConfig()
	}

	return &Snowfall{
		screen: screen,
		config: config,
		flakes: make([]*Snowflake, 0, 100),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化
func (s *Snowfall) Init() error {
	s.width, s.height = s.screen.Size()
	s.lastUpdate = time.Now()

	// 预生成一些雪花在屏幕各处
	initialFlakes := s.height / 2
	for i := 0; i < initialFlakes; i++ {
		flake := s.createFlake()
		flake.y = s.rand.Float64() * float64(s.height) // 随机 Y 位置
		s.flakes = append(s.flakes, flake)
	}

	return nil
}

// createFlake 创建新雪花
func (s *Snowfall) createFlake() *Snowflake {
	// 随机层次决定雪花外观
	layer := s.rand.Intn(3) // 0, 1, 2

	flake := &Snowflake{
		x:          s.rand.Float64() * float64(s.width),
		y:          0,
		char:       s.randomFlakeChar(layer),
		color:      s.getFlakeColor(layer),
		speed:      s.getFlakeSpeed(layer),
		swingPhase: s.rand.Float64() * 2 * math.Pi,
		swingAmp:   0.3 + s.rand.Float64()*0.7, // 0.3-1.0
		layer:      layer,
	}

	return flake
}

// randomFlakeChar 根据层次返回雪花字符
func (s *Snowfall) randomFlakeChar(layer int) rune {
	switch layer {
	case 0: // 远处 - 小字符
		chars := []rune{'·', '•', '.'}
		return chars[s.rand.Intn(len(chars))]
	case 1: // 中层 - 中等字符
		chars := []rune{'*', '❅', '·'}
		return chars[s.rand.Intn(len(chars))]
	case 2: // 近处 - 大字符
		chars := []rune{'❆', '❅', '*', '✻'}
		return chars[s.rand.Intn(len(chars))]
	}
	return '*'
}

// getFlakeColor 根据层次返回颜色
func (s *Snowfall) getFlakeColor(layer int) tcell.Color {
	switch layer {
	case 0: // 远处 - 暗色
		return tcell.ColorGray
	case 1: // 中层 - 白色
		return tcell.ColorWhite
	case 2: // 近处 - 亮白/淡蓝
		colors := []tcell.Color{tcell.ColorWhite, tcell.ColorLightBlue}
		return colors[s.rand.Intn(len(colors))]
	}
	return tcell.ColorWhite
}

// getFlakeSpeed 根据层次返回速度
func (s *Snowfall) getFlakeSpeed(layer int) float64 {
	switch layer {
	case 0: // 远处 - 慢
		return 0.5 + s.rand.Float64()*0.5 // 0.5-1.0
	case 1: // 中层 - 中速
		return 1.0 + s.rand.Float64()*0.5 // 1.0-1.5
	case 2: // 近处 - 快
		return 1.5 + s.rand.Float64()*0.5 // 1.5-2.0
	}
	return 1.0
}

// Update 更新雪花状态
func (s *Snowfall) Update(deltaTime float64) {
	// 生成新雪花
	s.generateNewFlakes()

	// 更新现有雪花
	alive := make([]*Snowflake, 0, len(s.flakes))
	for _, flake := range s.flakes {
		// 更新垂直位置
		flake.y += flake.speed * deltaTime * 10 // 10 是速度系数

		// 更新摆动相位
		flake.swingPhase += deltaTime * 2
		if flake.swingPhase > 2*math.Pi {
			flake.swingPhase -= 2 * math.Pi
		}

		// 计算水平摆动偏移
		swingOffset := math.Sin(flake.swingPhase) * flake.swingAmp

		// 更新水平位置
		flake.x += swingOffset * deltaTime * 5

		// 边界检查：水平环绕
		if flake.x < 0 {
			flake.x = float64(s.width)
		} else if flake.x >= float64(s.width) {
			flake.x = 0
		}

		// 垂直边界：到底部则删除
		if int(flake.y) < s.height {
			alive = append(alive, flake)
		}
	}

	s.flakes = alive
}

// generateNewFlakes 生成新雪花
func (s *Snowfall) generateNewFlakes() {
	if len(s.flakes) >= maxFlakes {
		return
	}

	var spawnCount int
	switch s.config.Density {
	case DensitySparse:
		if s.rand.Float64() < 0.3 { // 30% 概率生成 1-2 片
			spawnCount = 1 + s.rand.Intn(2)
		}
	case DensityMedium:
		spawnCount = 1 + s.rand.Intn(5) // 1-5 片
	case DensityDense:
		spawnCount = 3 + s.rand.Intn(8) // 3-10 片
	}

	for i := 0; i < spawnCount; i++ {
		if len(s.flakes) >= maxFlakes {
			break
		}
		s.flakes = append(s.flakes, s.createFlake())
	}
}

// Render 渲染雪花
func (s *Snowfall) Render() {
	s.screen.Clear()

	for _, flake := range s.flakes {
		x := int(flake.x)
		y := int(flake.y)

		if x >= 0 && x < s.width && y >= 0 && y < s.height {
			style := tcell.StyleDefault.Foreground(flake.color)
			s.screen.SetContent(x, y, flake.char, nil, style)
		}
	}

	s.screen.Show()
}

// Run 运行雪花特效
func (s *Snowfall) Run(quit <-chan struct{}) error {
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
func (s *Snowfall) Cleanup() error {
	return nil
}
