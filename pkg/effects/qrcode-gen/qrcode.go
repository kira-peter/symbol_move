package qrcodegen

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/skip2/go-qrcode"
)

// Config 二维码动画配置
type Config struct {
	ChangeInterval float64  // 二维码切换间隔（秒）
	Content        []string // 内容列表
	FPS            int      // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		ChangeInterval: 3.0, // 3秒切换
		Content: []string{
			"https://github.com/symbolmove/symbol_move",
			"SymbolMove - 符动世界",
			"终端特效展示",
			time.Now().Format("2006-01-02 15:04:05"),
		},
		FPS: 30,
	}
}

// QRCodeGen 二维码生成器特效
type QRCodeGen struct {
	screen      tcell.Screen
	config      *Config
	currentIdx  int
	timer       float64
	qrMatrix    [][]bool
	width       int
	height      int
	lastUpdate  time.Time
}

// New 创建二维码生成器特效实例
func New(screen tcell.Screen, config *Config) *QRCodeGen {
	if config == nil {
		config = DefaultConfig()
	}

	return &QRCodeGen{
		screen: screen,
		config: config,
	}
}

// Init 初始化二维码生成器
func (q *QRCodeGen) Init() error {
	q.width, q.height = q.screen.Size()
	q.currentIdx = 0
	q.timer = 0
	q.generateQR(q.config.Content[0])
	q.lastUpdate = time.Now()
	return nil
}

// generateQR 生成二维码
func (q *QRCodeGen) generateQR(content string) error {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}

	bitmap := qr.Bitmap()
	q.qrMatrix = bitmap
	return nil
}

// Update 更新二维码生成器状态
func (q *QRCodeGen) Update(deltaTime float64) {
	q.timer += deltaTime

	if q.timer >= q.config.ChangeInterval {
		q.timer = 0
		q.currentIdx = (q.currentIdx + 1) % len(q.config.Content)

		// 如果是时间戳，更新内容
		content := q.config.Content[q.currentIdx]
		if content == time.Now().Format("2006-01-02 15:04:05") {
			content = time.Now().Format("2006-01-02 15:04:05")
		}

		q.generateQR(content)
	}
}

// Render 渲染二维码
func (q *QRCodeGen) Render() {
	q.screen.Clear()

	if q.qrMatrix == nil || len(q.qrMatrix) == 0 {
		q.screen.Show()
		return
	}

	size := len(q.qrMatrix)

	// 计算缩放因子以适应屏幕
	maxSizeX := q.width / 2
	maxSizeY := q.height

	scale := 1
	if size > maxSizeX || size > maxSizeY {
		scaleX := maxSizeX / size
		scaleY := maxSizeY / size
		if scaleX < scaleY {
			scale = scaleX
		} else {
			scale = scaleY
		}
		if scale < 1 {
			scale = 1
		}
	} else {
		// 如果二维码太小，放大它
		scale = 2
	}

	// 计算居中位置
	displayW := size * scale * 2 // *2因为每个模块用2个字符宽度
	displayH := size * scale

	startX := (q.width - displayW) / 2
	startY := (q.height - displayH) / 2

	// 绘制二维码
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			module := q.qrMatrix[y][x]

			// 根据缩放因子绘制
			for sy := 0; sy < scale; sy++ {
				for sx := 0; sx < scale*2; sx++ {
					screenX := startX + x*scale*2 + sx
					screenY := startY + y*scale + sy

					if screenX >= 0 && screenX < q.width && screenY >= 0 && screenY < q.height {
						var style tcell.Style
						var char rune

						if module {
							// 黑色模块
							char = '█'
							style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
						} else {
							// 白色模块
							char = ' '
							style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorWhite)
						}

						q.screen.SetContent(screenX, screenY, char, nil, style)
					}
				}
			}
		}
	}

	// 显示当前内容文本
	content := q.config.Content[q.currentIdx]
	textY := startY + displayH + 2
	if textY < q.height {
		textX := (q.width - len(content)) / 2
		if textX < 0 {
			textX = 0
		}

		style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
		for i, ch := range content {
			if textX+i < q.width {
				q.screen.SetContent(textX+i, textY, ch, nil, style)
			}
		}
	}

	q.screen.Show()
}

// Run 运行二维码生成器特效
func (q *QRCodeGen) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(q.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(q.lastUpdate).Seconds()
			q.lastUpdate = now

			q.Update(deltaTime)
			q.Render()
		}
	}
}

// Cleanup 清理资源
func (q *QRCodeGen) Cleanup() error {
	return nil
}
