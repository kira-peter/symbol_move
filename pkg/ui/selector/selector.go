package selector

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// Selector 特效选择器
type Selector struct {
	screen       tcell.Screen
	effectList   []effects.Metadata
	selectedIdx  int
	width        int
	height       int
}

// New 创建新的选择器
func New(screen tcell.Screen) *Selector {
	return &Selector{
		screen:      screen,
		effectList:  effects.List(),
		selectedIdx: 0,
	}
}

// updateSize 更新终端尺寸
func (s *Selector) updateSize() {
	s.width, s.height = s.screen.Size()
}

// Render 渲染选择器界面
func (s *Selector) Render() {
	s.updateSize()
	s.screen.Clear()

	// 标题区域
	s.renderTitle()

	// 特效列表
	s.renderEffectList()

	// 描述区域
	s.renderDescription()

	// 提示区域
	s.renderHints()

	s.screen.Show()
}

// renderTitle 渲染标题
func (s *Selector) renderTitle() {
	title := "符动世界 (SymbolMove)"
	subtitle := "字符符号在动，创造世界"

	// 居中显示标题
	titleY := 2
	s.drawCenteredText(titleY, title, tcell.StyleDefault.
		Foreground(tcell.ColorLightGreen).
		Bold(true))

	s.drawCenteredText(titleY+1, subtitle, tcell.StyleDefault.
		Foreground(tcell.ColorGreen))

	// 分隔线
	s.drawHorizontalLine(titleY + 3)
}

// renderEffectList 渲染特效列表
func (s *Selector) renderEffectList() {
	startY := 6

	if len(s.effectList) == 0 {
		s.drawCenteredText(startY+2, "暂无可用特效", tcell.StyleDefault.
			Foreground(tcell.ColorGray))
		return
	}

	for i, metadata := range s.effectList {
		y := startY + i*2
		if y >= s.height-8 {
			break // 避免超出屏幕
		}

		// 序号
		indexText := fmt.Sprintf("%d.", i+1)

		// 特效名称
		nameText := metadata.Name

		// 完整文本
		text := fmt.Sprintf("  %s %s", indexText, nameText)

		// 选中状态
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
		if i == s.selectedIdx {
			// 高亮选中项
			text = fmt.Sprintf("► %s %s", indexText, nameText)
			style = tcell.StyleDefault.
				Foreground(tcell.ColorBlack).
				Background(tcell.ColorLightGreen).
				Bold(true)
		}

		s.drawText(s.width/2-20, y, text, style)
	}
}

// renderDescription 渲染描述区域
func (s *Selector) renderDescription() {
	if len(s.effectList) == 0 || s.selectedIdx >= len(s.effectList) {
		return
	}

	metadata := s.effectList[s.selectedIdx]
	descY := s.height - 6

	s.drawHorizontalLine(descY - 1)

	// 描述标题
	s.drawText(4, descY, "描述：", tcell.StyleDefault.
		Foreground(tcell.ColorYellow))

	// 描述内容
	desc := metadata.Description
	if len(desc) > s.width-12 {
		desc = desc[:s.width-15] + "..."
	}
	s.drawText(10, descY, desc, tcell.StyleDefault.
		Foreground(tcell.ColorWhite))
}

// renderHints 渲染操作提示
func (s *Selector) renderHints() {
	hintY := s.height - 2
	s.drawHorizontalLine(hintY - 1)

	hints := "↑↓: 选择  |  Enter: 确认  |  q/Ctrl+C: 退出"
	s.drawCenteredText(hintY, hints, tcell.StyleDefault.
		Foreground(tcell.ColorGray))
}

// drawText 在指定位置绘制文本
func (s *Selector) drawText(x, y int, text string, style tcell.Style) {
	if y < 0 || y >= s.height {
		return
	}

	for i, ch := range text {
		if x+i >= 0 && x+i < s.width {
			s.screen.SetContent(x+i, y, ch, nil, style)
		}
	}
}

// drawCenteredText 居中绘制文本
func (s *Selector) drawCenteredText(y int, text string, style tcell.Style) {
	x := (s.width - len(text)) / 2
	if x < 0 {
		x = 0
	}
	s.drawText(x, y, text, style)
}

// drawHorizontalLine 绘制水平分隔线
func (s *Selector) drawHorizontalLine(y int) {
	if y < 0 || y >= s.height {
		return
	}

	line := strings.Repeat("─", s.width)
	s.drawText(0, y, line, tcell.StyleDefault.
		Foreground(tcell.ColorGray))
}

// HandleKey 处理键盘事件
func (s *Selector) HandleKey(event *tcell.EventKey) int {
	if len(s.effectList) == 0 {
		return -1
	}

	switch event.Key() {
	case tcell.KeyUp:
		s.MoveUp()
	case tcell.KeyDown:
		s.MoveDown()
	case tcell.KeyEnter:
		return s.selectedIdx
	case tcell.KeyRune:
		switch event.Rune() {
		case 'k', 'K':
			s.MoveUp()
		case 'j', 'J':
			s.MoveDown()
		case 'q', 'Q':
			return -2 // 退出信号
		default:
			// 数字快捷键
			if event.Rune() >= '1' && event.Rune() <= '9' {
				idx := int(event.Rune() - '1')
				if idx < len(s.effectList) {
					s.selectedIdx = idx
					return s.selectedIdx
				}
			}
		}
	}

	return -1 // 无操作
}

// MoveUp 向上移动选择
func (s *Selector) MoveUp() {
	if len(s.effectList) == 0 {
		return
	}

	s.selectedIdx--
	if s.selectedIdx < 0 {
		s.selectedIdx = len(s.effectList) - 1 // 循环到末尾
	}
}

// MoveDown 向下移动选择
func (s *Selector) MoveDown() {
	if len(s.effectList) == 0 {
		return
	}

	s.selectedIdx++
	if s.selectedIdx >= len(s.effectList) {
		s.selectedIdx = 0 // 循环到开头
	}
}

// GetSelected 获取当前选中的特效元数据
func (s *Selector) GetSelected() (effects.Metadata, bool) {
	if len(s.effectList) == 0 || s.selectedIdx >= len(s.effectList) {
		return effects.Metadata{}, false
	}

	return s.effectList[s.selectedIdx], true
}

// Refresh 刷新特效列表
func (s *Selector) Refresh() {
	s.effectList = effects.List()
	if s.selectedIdx >= len(s.effectList) {
		s.selectedIdx = 0
	}
}
