package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	matrixrain "github.com/symbolmove/symbol_move/pkg/effects/matrix-rain"
)

func main() {
	// 命令行参数
	var (
		speed    string
		density  string
		charset  string
		fps      int
		help     bool
	)

	flag.StringVar(&speed, "speed", "medium", "下落速度: slow, medium, fast")
	flag.StringVar(&density, "density", "medium", "字符雨密度: sparse, medium, dense")
	flag.StringVar(&charset, "charset", "mixed", "字符集: digits, letters, katakana, mixed")
	flag.IntVar(&fps, "fps", 30, "帧率 (默认 30)")
	flag.BoolVar(&help, "help", false, "显示帮助信息")

	flag.Parse()

	if help {
		printHelp()
		return
	}

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

	// 确保退出时恢复终端
	defer screen.Fini()

	// 创建配置
	config := matrixrain.DefaultConfig()
	config.FPS = fps

	// 解析速度
	switch speed {
	case "slow":
		config.Speed = matrixrain.SpeedSlow
	case "medium":
		config.Speed = matrixrain.SpeedMedium
	case "fast":
		config.Speed = matrixrain.SpeedFast
	default:
		fmt.Fprintf(os.Stderr, "未知速度: %s (使用 slow, medium, fast)\n", speed)
		os.Exit(1)
	}

	// 解析密度
	switch density {
	case "sparse":
		config.Density = matrixrain.DensitySparse
	case "medium":
		config.Density = matrixrain.DensityMedium
	case "dense":
		config.Density = matrixrain.DensityDense
	default:
		fmt.Fprintf(os.Stderr, "未知密度: %s (使用 sparse, medium, dense)\n", density)
		os.Exit(1)
	}

	// 解析字符集
	switch charset {
	case "digits":
		config.CharSet = matrixrain.CharSetDigits
	case "letters":
		config.CharSet = matrixrain.CharSetLetters
	case "katakana":
		config.CharSet = matrixrain.CharSetKatakana
	case "mixed":
		config.CharSet = matrixrain.CharSetMixed
	default:
		fmt.Fprintf(os.Stderr, "未知字符集: %s (使用 digits, letters, katakana, mixed)\n", charset)
		os.Exit(1)
	}

	// 创建字符雨效果
	rain := matrixrain.New(screen, config)

	// 主循环
	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				// 按 ESC, Ctrl+C 或 q 退出
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
					close(quit)
					return
				}
			case *tcell.EventResize:
				// 处理终端大小调整
				rain.Resize()
				screen.Sync()
			}
		}
	}()

	// 帧率控制
	ticker := time.NewTicker(time.Second / time.Duration(fps))
	defer ticker.Stop()

	// 渲染循环
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			rain.Update()
			rain.Render()
		}
	}
}

func printHelp() {
	fmt.Println("符动世界 - 矩阵字符雨效果")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  matrix-rain [选项]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  -speed string")
	fmt.Println("        下落速度: slow, medium, fast (默认 medium)")
	fmt.Println("  -density string")
	fmt.Println("        字符雨密度: sparse, medium, dense (默认 medium)")
	fmt.Println("  -charset string")
	fmt.Println("        字符集: digits, letters, katakana, mixed (默认 mixed)")
	fmt.Println("  -fps int")
	fmt.Println("        帧率 (默认 30)")
	fmt.Println("  -help")
	fmt.Println("        显示此帮助信息")
	fmt.Println()
	fmt.Println("控制:")
	fmt.Println("  ESC, Ctrl+C, q  - 退出程序")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  matrix-rain")
	fmt.Println("  matrix-rain -speed fast -density dense")
	fmt.Println("  matrix-rain -charset katakana -speed slow")
}
