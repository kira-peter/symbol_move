package matrixrain

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// CharSet 定义字符集类型
type CharSet string

const (
	CharSetDigits   CharSet = "digits"   // 数字 0-9
	CharSetLetters  CharSet = "letters"  // 字母 A-Z, a-z
	CharSetKatakana CharSet = "katakana" // 日文片假名
	CharSetMixed    CharSet = "mixed"    // 混合
	CharSetCustom   CharSet = "custom"   // 自定义
)

// Speed 定义速度档位
type Speed int

const (
	SpeedSlow Speed = iota
	SpeedMedium
	SpeedFast
)

// Density 定义密度档位
type Density int

const (
	DensitySparse Density = iota
	DensityMedium
	DensityDense
)

// Config 字符雨配置
type Config struct {
	CharSet       CharSet   // 字符集类型
	CustomChars   []rune    // 自定义字符集
	Speed         Speed     // 下落速度
	Density       Density   // 字符雨密度
	FPS           int       // 帧率
	TrailLength   int       // 字符流尾迹长度
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		CharSet:     CharSetMixed,
		Speed:       SpeedMedium,
		Density:     DensityMedium,
		FPS:         30,
		TrailLength: 15,
	}
}

// RainDrop 表示一条字符流
type RainDrop struct {
	X      int     // X 坐标（列）
	Y      float64 // Y 坐标（行，使用浮点数实现平滑移动）
	Speed  float64 // 下落速度
	Length int     // 字符流长度
	Chars  []rune  // 字符内容
}

// Rain 字符雨效果管理器
type Rain struct {
	screen      tcell.Screen
	config      *Config
	drops       []*RainDrop
	width       int
	height      int
	charPool    []rune
	rand        *rand.Rand
	lastUpdate  time.Time
}

// New 创建新的字符雨效果
func New(screen tcell.Screen, config *Config) *Rain {
	if config == nil {
		config = DefaultConfig()
	}

	r := &Rain{
		screen:     screen,
		config:     config,
		drops:      make([]*RainDrop, 0),
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		lastUpdate: time.Now(),
	}

	r.updateSize()
	r.initCharPool()
	r.initDrops()

	return r
}

// updateSize 更新终端尺寸
func (r *Rain) updateSize() {
	r.width, r.height = r.screen.Size()
}

// initCharPool 初始化字符池
func (r *Rain) initCharPool() {
	switch r.config.CharSet {
	case CharSetDigits:
		r.charPool = []rune("0123456789")
	case CharSetLetters:
		r.charPool = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	case CharSetKatakana:
		// 日文片假名范围
		katakana := make([]rune, 0)
		for r := rune(0x30A0); r <= rune(0x30FF); r++ {
			katakana = append(katakana, r)
		}
		r.charPool = katakana
	case CharSetMixed:
		// 混合：数字 + 字母 + 部分片假名
		mixed := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		for c := rune(0x30A0); c <= rune(0x30CF); c++ {
			mixed = append(mixed, c)
		}
		r.charPool = mixed
	case CharSetCustom:
		if len(r.config.CustomChars) > 0 {
			r.charPool = r.config.CustomChars
		} else {
			// 如果自定义字符集为空，回退到数字
			r.charPool = []rune("0123456789")
		}
	default:
		r.charPool = []rune("0123456789")
	}
}

// initDrops 初始化字符流
func (r *Rain) initDrops() {
	// 根据密度计算字符流数量
	var dropCount int
	switch r.config.Density {
	case DensitySparse:
		dropCount = r.width / 4
	case DensityMedium:
		dropCount = r.width / 2
	case DensityDense:
		dropCount = r.width
	}

	for i := 0; i < dropCount; i++ {
		r.drops = append(r.drops, r.createRandomDrop())
	}
}

// createRandomDrop 创建随机字符流
func (r *Rain) createRandomDrop() *RainDrop {
	length := r.rand.Intn(r.config.TrailLength) + 5
	chars := make([]rune, length)
	for i := range chars {
		chars[i] = r.randomChar()
	}

	// 计算速度
	var speed float64
	switch r.config.Speed {
	case SpeedSlow:
		speed = 0.2 + r.rand.Float64()*0.2 // 0.2-0.4
	case SpeedMedium:
		speed = 0.4 + r.rand.Float64()*0.3 // 0.4-0.7
	case SpeedFast:
		speed = 0.7 + r.rand.Float64()*0.5 // 0.7-1.2
	}

	return &RainDrop{
		X:      r.rand.Intn(r.width),
		Y:      -float64(length), // 从屏幕上方开始
		Speed:  speed,
		Length: length,
		Chars:  chars,
	}
}

// randomChar 随机选择一个字符
func (r *Rain) randomChar() rune {
	if len(r.charPool) == 0 {
		return '0'
	}
	return r.charPool[r.rand.Intn(len(r.charPool))]
}

// Update 更新字符雨状态
func (r *Rain) Update() {
	now := time.Now()
	dt := now.Sub(r.lastUpdate).Seconds()
	r.lastUpdate = now

	// 更新所有字符流
	for i := 0; i < len(r.drops); i++ {
		drop := r.drops[i]

		// 更新位置
		drop.Y += drop.Speed * dt * float64(r.config.FPS)

		// 随机改变顶部字符
		if r.rand.Float64() < 0.1 {
			drop.Chars[0] = r.randomChar()
		}

		// 如果字符流完全离开屏幕，重置
		if int(drop.Y) > r.height {
			r.drops[i] = r.createRandomDrop()
		}
	}
}

// Render 渲染字符雨到屏幕
func (r *Rain) Render() {
	r.screen.Clear()

	for _, drop := range r.drops {
		r.renderDrop(drop)
	}

	r.screen.Show()
}

// renderDrop 渲染单个字符流
func (r *Rain) renderDrop(drop *RainDrop) {
	for i := 0; i < drop.Length; i++ {
		y := int(drop.Y) - i

		// 只渲染在屏幕范围内的字符
		if y < 0 || y >= r.height {
			continue
		}
		if drop.X < 0 || drop.X >= r.width {
			continue
		}

		char := drop.Chars[i%len(drop.Chars)]
		style := r.getCharStyle(i, drop.Length)

		r.screen.SetContent(drop.X, y, char, nil, style)
	}
}

// getCharStyle 根据位置返回字符样式（实现颜色渐变）
func (r *Rain) getCharStyle(index, length int) tcell.Style {
	// index 0 是最新（顶部），index length-1 是最旧（底部）
	ratio := float64(index) / float64(length)

	var fg tcell.Color

	if ratio < 0.1 {
		// 顶部：亮白色
		fg = tcell.ColorWhite
	} else if ratio < 0.3 {
		// 接近顶部：亮绿色
		fg = tcell.ColorLightGreen
	} else if ratio < 0.7 {
		// 中间：标准绿色
		fg = tcell.ColorGreen
	} else {
		// 底部：暗绿色
		fg = tcell.ColorDarkGreen
	}

	return tcell.StyleDefault.Foreground(fg).Background(tcell.ColorBlack)
}

// Resize 处理终端大小调整
func (r *Rain) Resize() {
	oldWidth := r.width
	r.updateSize()

	// 如果宽度改变，调整字符流数量
	if oldWidth != r.width {
		r.drops = r.drops[:0]
		r.initDrops()
	}
}
