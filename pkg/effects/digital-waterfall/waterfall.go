package digitalwaterfall

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 数字瀑布配置
type Config struct {
	MinSpeed  float64 // 最小流动速度
	MaxSpeed  float64 // 最大流动速度
	MinLength int     // 最小流长度
	MaxLength int     // 最大流长度
	FPS       int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		MinSpeed:  0.5,
		MaxSpeed:  2.0,
		MinLength: 5,
		MaxLength: 15,
		FPS:       30,
	}
}

// Column 单个数字流列
type Column struct {
	x          int       // 列的 x 坐标
	y          float64   // 当前头部的 y 坐标（浮点数用于平滑移动）
	speed      float64   // 流动速度（单位：格/帧）
	length     int       // 数字流长度
	digits     []rune    // 数字序列
	spawnDelay int       // 生成延迟（帧数）
}

// DigitalWaterfall 数字瀑布特效
type DigitalWaterfall struct {
	screen  tcell.Screen
	config  *Config
	columns []*Column
	width   int
	height  int
	rand    *rand.Rand
}

// New 创建数字瀑布特效实例
func New(screen tcell.Screen, config *Config) *DigitalWaterfall {
	if config == nil {
		config = DefaultConfig()
	}

	return &DigitalWaterfall{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化数字瀑布
func (d *DigitalWaterfall) Init() error {
	d.width, d.height = d.screen.Size()
	d.initColumns()
	return nil
}

// initColumns 初始化所有列
func (d *DigitalWaterfall) initColumns() {
	d.columns = make([]*Column, d.width)

	for x := 0; x < d.width; x++ {
		d.columns[x] = d.createColumn(x)
		// 随机初始位置，避免所有列同时开始
		d.columns[x].y = -float64(d.rand.Intn(d.height))
	}
}

// createColumn 创建新的数字流列
func (d *DigitalWaterfall) createColumn(x int) *Column {
	length := d.config.MinLength + d.rand.Intn(d.config.MaxLength-d.config.MinLength+1)
	speed := d.config.MinSpeed + d.rand.Float64()*(d.config.MaxSpeed-d.config.MinSpeed)

	column := &Column{
		x:      x,
		y:      -float64(length), // 从屏幕上方开始
		speed:  speed,
		length: length,
		digits: make([]rune, length),
	}

	// 生成随机数字序列
	for i := 0; i < length; i++ {
		column.digits[i] = d.randomDigit()
	}

	return column
}

// randomDigit 生成随机数字字符
func (d *DigitalWaterfall) randomDigit() rune {
	return rune('0' + d.rand.Intn(10))
}

// Update 更新数字瀑布状态
func (d *DigitalWaterfall) Update() {
	for _, column := range d.columns {
		// 移动列
		column.y += column.speed

		// 随机改变头部数字（制造变化效果）
		if d.rand.Float64() < 0.3 {
			column.digits[0] = d.randomDigit()
		}

		// 如果列完全离开屏幕，重置
		if column.y-float64(column.length) > float64(d.height) {
			*column = *d.createColumn(column.x)
		}
	}
}

// Render 渲染数字瀑布
func (d *DigitalWaterfall) Render() {
	d.screen.Clear()

	for _, column := range d.columns {
		d.renderColumn(column)
	}

	d.screen.Show()
}

// renderColumn 渲染单个列
func (d *DigitalWaterfall) renderColumn(column *Column) {
	for i := 0; i < column.length; i++ {
		y := int(column.y) - i

		// 只渲染在屏幕内的数字
		if y >= 0 && y < d.height {
			// 计算亮度（头部最亮，尾部最暗）
			brightness := float64(column.length-i) / float64(column.length)
			style := d.getDigitStyle(brightness, i == 0)

			d.screen.SetContent(column.x, y, column.digits[i], nil, style)
		}
	}
}

// getDigitStyle 根据亮度获取数字样式
func (d *DigitalWaterfall) getDigitStyle(brightness float64, isHead bool) tcell.Style {
	var color tcell.Color

	// 头部偶尔白色闪光
	if isHead && d.rand.Float64() < 0.05 {
		color = tcell.ColorWhite
	} else if brightness > 0.8 {
		color = tcell.ColorLime // 最亮绿色
	} else if brightness > 0.5 {
		color = tcell.ColorGreen // 中绿色
	} else if brightness > 0.3 {
		color = tcell.ColorDarkGreen // 暗绿色
	} else {
		color = tcell.ColorGray // 最暗
	}

	style := tcell.StyleDefault.Foreground(color)

	// 头部加粗
	if isHead {
		style = style.Bold(true)
	}

	return style
}

// Run 运行数字瀑布特效
func (d *DigitalWaterfall) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(d.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			d.Update()
			d.Render()
		}
	}
}

// Cleanup 清理资源
func (d *DigitalWaterfall) Cleanup() error {
	return nil
}
