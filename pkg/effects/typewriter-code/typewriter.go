package typewritercode

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 打字机代码雨配置
type Config struct {
	TypingSpeed  float64 // 字符/秒
	LineInterval float64 // 新行间隔（秒）
	FPS          int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		TypingSpeed:  20.0, // 20字符/秒
		LineInterval: 0.5,  // 0.5秒一行
		FPS:          30,
	}
}

// CodeLine 代码行
type CodeLine struct {
	text     string
	x, y     int
	progress float64 // 当前显示到第几个字符（浮点数以支持平滑速度）
	speed    float64 // 打字速度倍数
	finished bool    // 是否打字完成
}

// codeTemplates 代码模板库
var codeTemplates = []string{
	"func main() {",
	"    for i := 0; i < 10; i++ {",
	"        fmt.Println(i)",
	"    }",
	"}",
	"package main",
	"import \"fmt\"",
	"class Example:",
	"    def __init__(self):",
	"        self.value = 42",
	"    def process(self):",
	"        return self.value * 2",
	"const express = require('express');",
	"const app = express();",
	"app.get('/', (req, res) => {",
	"    res.send('Hello World!');",
	"});",
	"if err != nil {",
	"    return err",
	"for _, item := range items {",
	"    process(item)",
	"let result = data.filter(x => x > 0)",
	"async function fetchData() {",
	"    const response = await fetch(url);",
	"    return response.json();",
}

// Typewriter 打字机代码雨特效
type Typewriter struct {
	screen            tcell.Screen
	config            *Config
	lines             []*CodeLine
	width             int
	height            int
	timeSinceNewLine  float64
	lastUpdate        time.Time
	rand              *rand.Rand
}

// New 创建打字机代码雨特效实例
func New(screen tcell.Screen, config *Config) *Typewriter {
	if config == nil {
		config = DefaultConfig()
	}

	return &Typewriter{
		screen: screen,
		config: config,
		lines:  make([]*CodeLine, 0),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化打字机代码雨
func (t *Typewriter) Init() error {
	t.width, t.height = t.screen.Size()
	t.timeSinceNewLine = 0
	t.lastUpdate = time.Now()
	return nil
}

// addNewLine 添加新代码行
func (t *Typewriter) addNewLine() {
	// 移除超出屏幕的行
	activeLines := make([]*CodeLine, 0)
	for _, line := range t.lines {
		if line.y < t.height {
			activeLines = append(activeLines, line)
		}
	}

	// 所有现有行下移一行
	for _, line := range activeLines {
		line.y++
	}

	// 添加新行
	line := &CodeLine{
		text:     codeTemplates[t.rand.Intn(len(codeTemplates))],
		x:        0,
		y:        0,
		progress: 0,
		speed:    0.8 + t.rand.Float64()*0.4, // 0.8-1.2倍速
		finished: false,
	}

	activeLines = append([]*CodeLine{line}, activeLines...)
	t.lines = activeLines
}

// getColorForChar 获取字符颜色（简单语法高亮）
func (t *Typewriter) getColorForChar(text string, pos int) tcell.Color {
	// 简单的关键字匹配
	keywords := []string{"func", "for", "if", "return", "class", "def", "const", "let", "async", "import", "package"}

	for _, keyword := range keywords {
		if pos+len(keyword) <= len(text) {
			if text[pos:pos+len(keyword)] == keyword {
				return tcell.ColorYellow
			}
		}
	}

	ch := rune(text[pos])
	switch {
	case ch == '"' || ch == '\'':
		return tcell.ColorGreen
	case ch >= '0' && ch <= '9':
		return tcell.ColorLightCyan
	case ch == '(' || ch == ')' || ch == '{' || ch == '}' || ch == '[' || ch == ']':
		return tcell.ColorPurple
	default:
		return tcell.ColorWhite
	}
}

// Update 更新打字机代码雨状态
func (t *Typewriter) Update(deltaTime float64) {
	// 定期添加新行
	t.timeSinceNewLine += deltaTime
	if t.timeSinceNewLine >= t.config.LineInterval {
		t.addNewLine()
		t.timeSinceNewLine = 0
	}

	// 更新每行的打字进度
	for _, line := range t.lines {
		if !line.finished {
			line.progress += deltaTime * t.config.TypingSpeed * line.speed
			if line.progress >= float64(len(line.text)) {
				line.progress = float64(len(line.text))
				line.finished = true
			}
		}
	}
}

// Render 渲染打字机代码雨
func (t *Typewriter) Render() {
	t.screen.Clear()

	for _, line := range t.lines {
		if line.y < 0 || line.y >= t.height {
			continue
		}

		// 只显示已打字的部分
		displayLen := int(line.progress)
		if displayLen > len(line.text) {
			displayLen = len(line.text)
		}

		visibleText := line.text[:displayLen]

		// 渲染每个字符
		x := line.x
		for i, ch := range visibleText {
			if x >= t.width {
				break
			}

			color := t.getColorForChar(line.text, i)
			style := tcell.StyleDefault.Foreground(color)

			// 正在打字的字符闪烁
			if i == displayLen-1 && !line.finished {
				style = style.Bold(true)
			}

			t.screen.SetContent(x, line.y, ch, nil, style)
			x++
		}

		// 显示光标（正在打字的位置）
		if !line.finished && x < t.width {
			style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Bold(true)
			t.screen.SetContent(x, line.y, '▌', nil, style)
		}
	}

	t.screen.Show()
}

// Run 运行打字机代码雨特效
func (t *Typewriter) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(t.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(t.lastUpdate).Seconds()
			t.lastUpdate = now

			t.Update(deltaTime)
			t.Render()
		}
	}
}

// Cleanup 清理资源
func (t *Typewriter) Cleanup() error {
	return nil
}
