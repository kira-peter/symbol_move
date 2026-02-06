package audiovisualizer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

type AudioVisualizerEffect struct {
	visualizer *AudioVisualizer
	config     *Config
}

func NewEffect() effects.Effect {
	return &AudioVisualizerEffect{
		config: DefaultConfig(),
	}
}

func (e *AudioVisualizerEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "audio-visualizer",
		Name:        "音频可视化",
		Description: "模拟音频波形展示，使用随机或正弦波数据，频谱柱状图效果",
		LongDescription: `
音频可视化特效模拟音频频谱显示效果。

特点：
- 频谱柱状图
- 平滑动画过渡
- 动态颜色渐变
- 模拟音频数据
- 30 FPS 流畅运行

完美用于：
- 音乐播放器效果
- 音频展示
- 可视化演示
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"音频", "可视化", "频谱", "动画"},
	}
}

func (e *AudioVisualizerEffect) Init(screen tcell.Screen) error {
	e.visualizer = New(screen, e.config)
	return e.visualizer.Init()
}

func (e *AudioVisualizerEffect) Run(quit <-chan struct{}) error {
	return e.visualizer.Run(quit)
}

func (e *AudioVisualizerEffect) Cleanup() error {
	if e.visualizer != nil {
		return e.visualizer.Cleanup()
	}
	return nil
}
