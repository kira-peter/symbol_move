package waterripple

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 水波涟漪配置
type Config struct {
	DropInterval float64 // 水滴间隔（秒）
	WaveSpeed    float64 // 波速
	Damping      float64 // 衰减系数
	FPS          int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		DropInterval: 2.0,  // 2秒一滴
		WaveSpeed:    3.0,  // 波速
		Damping:      0.3,  // 衰减系数
		FPS:          30,
	}
}

// Drop 水滴
type Drop struct {
	x, y      int     // 水滴中心
	time      float64 // 已存在时间
	maxRadius float64 // 最大半径
}

// WaterRipple 水波涟漪特效
type WaterRipple struct {
	screen           tcell.Screen
	config           *Config
	drops            []*Drop
	width            int
	height           int
	timeSinceNewDrop float64
	lastUpdate       time.Time
	rand             *rand.Rand
}

// New 创建水波涟漪特效实例
func New(screen tcell.Screen, config *Config) *WaterRipple {
	if config == nil {
		config = DefaultConfig()
	}

	return &WaterRipple{
		screen: screen,
		config: config,
		drops:  make([]*Drop, 0),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化水波涟漪
func (w *WaterRipple) Init() error {
	w.width, w.height = w.screen.Size()
	w.timeSinceNewDrop = 0
	w.lastUpdate = time.Now()
	return nil
}

// addDrop 添加新水滴
func (w *WaterRipple) addDrop() {
	drop := &Drop{
		x:         w.rand.Intn(w.width),
		y:         w.rand.Intn(w.height),
		time:      0,
		maxRadius: 20.0 + w.rand.Float64()*10.0,
	}
	w.drops = append(w.drops, drop)
}

// Update 更新水波涟漪状态
func (w *WaterRipple) Update(deltaTime float64) {
	// 添加新水滴
	w.timeSinceNewDrop += deltaTime
	if w.timeSinceNewDrop >= w.config.DropInterval {
		w.addDrop()
		w.timeSinceNewDrop = 0
	}

	// 更新现有水滴
	activeDrops := make([]*Drop, 0)
	for _, drop := range w.drops {
		drop.time += deltaTime

		// 移除过期的水滴
		currentRadius := w.config.WaveSpeed * drop.time
		if currentRadius < drop.maxRadius {
			activeDrops = append(activeDrops, drop)
		}
	}
	w.drops = activeDrops
}

// Render 渲染水波涟漪
func (w *WaterRipple) Render() {
	w.screen.Clear()

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			// 叠加所有水滴的波动
			totalAmplitude := 0.0

			for _, drop := range w.drops {
				// 距离水滴中心的距离
				dx := float64(x - drop.x)
				dy := float64(y - drop.y)
				r := math.Sqrt(dx*dx + dy*dy)

				currentRadius := w.config.WaveSpeed * drop.time

				// 只在波前附近有显著振幅
				if math.Abs(r-currentRadius) < 3.0 && r > 0 {
					// 波动方程
					k := 0.5 // 波数
					omega := w.config.WaveSpeed
					amplitude := math.Sin(k*r-omega*drop.time) / (1 + r/10)

					// 衰减
					amplitude *= math.Exp(-w.config.Damping * drop.time)
					totalAmplitude += amplitude
				}
			}

			// 根据振幅选择字符和颜色
			if math.Abs(totalAmplitude) > 0.15 {
				var char rune
				var color tcell.Color

				if totalAmplitude > 0.5 {
					char = '○'
					color = tcell.ColorWhite
				} else if totalAmplitude > 0.3 {
					char = '◯'
					color = tcell.ColorLightCyan
				} else if totalAmplitude < -0.3 {
					char = '·'
					color = tcell.ColorBlue
				} else {
					char = '~'
					color = tcell.ColorLightBlue
				}

				style := tcell.StyleDefault.Foreground(color)
				w.screen.SetContent(x, y, char, nil, style)
			}
		}
	}

	w.screen.Show()
}

// Run 运行水波涟漪特效
func (w *WaterRipple) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(w.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(w.lastUpdate).Seconds()
			w.lastUpdate = now

			w.Update(deltaTime)
			w.Render()
		}
	}
}

// Cleanup 清理资源
func (w *WaterRipple) Cleanup() error {
	return nil
}
