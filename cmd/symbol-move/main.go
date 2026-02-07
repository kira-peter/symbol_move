package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
	_ "github.com/symbolmove/symbol_move/pkg/effects/audio-visualizer"  // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/big-clock"         // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/digital-waterfall" // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/dna-helix"         // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/fire-effect"       // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/fireworks"         // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/game-of-life"      // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/heartbeat"         // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/matrix-rain"       // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/matrix-tunnel"     // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/maze-generator"    // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/ocean-wave"        // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/particle-burst"    // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/plasma"            // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/qrcode-gen"        // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/rainbow-wave"      // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/snake-ai"          // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/snowfall"          // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/starry-sky"        // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/tetris-auto"       // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/typewriter-code"   // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/water-ripple"      // 自动注册
	_ "github.com/symbolmove/symbol_move/pkg/effects/wave-text"         // 自动注册
	"github.com/symbolmove/symbol_move/pkg/i18n"
	"github.com/symbolmove/symbol_move/pkg/ui/selector"
)

func main() {
	// 加载用户语言配置
	mgr := i18n.GetManager()
	mgr.LoadConfig() // 忽略错误，使用默认值

	// 初始化终端
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "初始化终端失败: %v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "初始化屏幕失败: %v\n", err)
		os.Exit(1)
	}

	defer screen.Fini()

	// 运行主循环
	if err := runMainLoop(screen); err != nil {
		screen.Fini()
		fmt.Fprintf(os.Stderr, "运行错误: %v\n", err)
		os.Exit(1)
	}
}

// runMainLoop 主循环 - 选择器和特效之间的状态机
func runMainLoop(screen tcell.Screen) error {
	sel := selector.New(screen)

	for {
		// 显示选择器界面
		sel.Refresh()
		sel.Render()

		// 等待用户选择
		_, quit := waitForSelection(screen, sel)

		if quit {
			// 用户选择退出
			return nil
		}

		// 获取选中的特效
		metadata, ok := sel.GetSelected()
		if !ok {
			continue
		}

		// 运行特效
		if err := runEffect(screen, metadata.ID); err != nil {
			// 显示错误（简单处理）
			showError(screen, fmt.Sprintf("特效运行错误: %v", err))
			continue
		}

		// 特效结束后，循环回到选择器
		sel.Render()
	}
}

// waitForSelection 等待用户选择
func waitForSelection(screen tcell.Screen, sel *selector.Selector) (int, bool) {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return -1, true // 退出
			}

			result := sel.HandleKey(ev)
			if result == -2 {
				return -1, true // 用户按 q 退出
			}
			if result == -3 {
				// 语言切换，重新渲染
				sel.Render()
				continue
			}
			if result >= 0 {
				return result, false // 用户选择了特效
			}

			// 重新渲染（选择可能改变）
			sel.Render()

		case *tcell.EventResize:
			sel.Render()
		}
	}
}

// runEffect 运行指定的特效
func runEffect(screen tcell.Screen, effectID string) error {
	// 获取特效工厂
	factory, err := effects.Get(effectID)
	if err != nil {
		return err
	}

	// 创建特效实例
	effect := factory()

	// 初始化特效
	if err := effect.Init(screen); err != nil {
		return fmt.Errorf("初始化失败: %w", err)
	}

	// 清理资源
	defer effect.Cleanup()

	// 创建退出通道
	quit := make(chan struct{})

	// 启动键盘监听协程
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				// ESC 键返回主界面
				if ev.Key() == tcell.KeyEscape {
					close(quit)
					return
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	// 运行特效
	return effect.Run(quit)
}

// showError 显示错误信息
func showError(screen tcell.Screen, message string) {
	screen.Clear()

	width, height := screen.Size()
	y := height / 2

	// 错误标题
	title := "错误"
	x := (width - len(title)) / 2
	style := tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)
	for i, ch := range title {
		screen.SetContent(x+i, y-1, ch, nil, style)
	}

	// 错误信息
	x = (width - len(message)) / 2
	if x < 0 {
		x = 0
	}
	style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	for i, ch := range message {
		if x+i < width {
			screen.SetContent(x+i, y+1, ch, nil, style)
		}
	}

	// 提示
	hint := "按任意键继续..."
	x = (width - len(hint)) / 2
	style = tcell.StyleDefault.Foreground(tcell.ColorGray)
	for i, ch := range hint {
		screen.SetContent(x+i, y+3, ch, nil, style)
	}

	screen.Show()

	// 等待按键
	screen.PollEvent()
}
