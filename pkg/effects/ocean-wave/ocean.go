package oceanwave

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 字符海浪配置
type Config struct {
	WaveSpeed  float64 // 波浪速度
	WaveHeight float64 // 波浪高度
	NumLayers  int     // 波浪层数
	FPS        int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		WaveSpeed:  2.0,  // 波浪速度
		WaveHeight: 4.0,  // 波浪高度
		NumLayers:  3,    // 3层波浪叠加
		FPS:        30,
	}
}

// OceanWave 字符海浪特效
type OceanWave struct {
	screen     tcell.Screen
	config     *Config
	phase      float64
	width      int
	height     int
	lastUpdate time.Time
	rand       *rand.Rand
}

// New 创建字符海浪特效实例
func New(screen tcell.Screen, config *Config) *OceanWave {
	if config == nil {
		config = DefaultConfig()
	}

	return &OceanWave{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化字符海浪
func (o *OceanWave) Init() error {
	o.width, o.height = o.screen.Size()
	o.phase = 0
	o.lastUpdate = time.Now()
	return nil
}

// Update 更新字符海浪状态
func (o *OceanWave) Update(deltaTime float64) {
	o.phase += o.config.WaveSpeed * deltaTime
	// 保持相位在合理范围内
	if o.phase > 2*math.Pi*100 {
		o.phase -= 2 * math.Pi * 100
	}
}

// Render 渲染字符海浪
func (o *OceanWave) Render() {
	o.screen.Clear()

	baseY := o.height * 2 / 3 // 海平面位置

	for x := 0; x < o.width; x++ {
		// 多层波浪叠加
		waveY := 0.0
		for i := 0; i < o.config.NumLayers; i++ {
			freq := float64(i+1) * 0.3
			amp := o.config.WaveHeight / float64(i+1)
			waveY += amp * math.Sin(o.phase+float64(x)*0.05*freq)
		}

		finalY := baseY + int(waveY)

		// 绘制海水和浪花
		for y := finalY; y < o.height; y++ {
			if y < 0 || y >= o.height {
				continue
			}

			depth := y - finalY
			var char rune
			var color tcell.Color

			if depth == 0 {
				// 浪花顶部
				foamChars := []rune{'~', '≈', '∿'}
				char = foamChars[o.rand.Intn(len(foamChars))]
				color = tcell.ColorWhite
			} else if depth < 3 {
				// 浅水区
				char = '≈'
				if depth == 1 {
					color = tcell.ColorLightCyan
				} else {
					color = tcell.ColorLightBlue
				}
			} else if depth < 6 {
				// 中等深度
				char = '~'
				color = tcell.ColorBlue
			} else {
				// 深海
				char = '~'
				color = tcell.ColorDarkBlue
			}

			style := tcell.StyleDefault.Foreground(color)
			o.screen.SetContent(x, y, char, nil, style)
		}
	}

	o.screen.Show()
}

// Run 运行字符海浪特效
func (o *OceanWave) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(o.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(o.lastUpdate).Seconds()
			o.lastUpdate = now

			o.Update(deltaTime)
			o.Render()
		}
	}
}

// Cleanup 清理资源
func (o *OceanWave) Cleanup() error {
	return nil
}
