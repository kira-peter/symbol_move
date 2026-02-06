package heartbeat

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// HeartbeatEffect 心跳律动特效
type HeartbeatEffect struct {
	heartbeat *Heartbeat
	config    *Config
}

// NewEffect 创建心跳律动特效实例
func NewEffect() effects.Effect {
	return &HeartbeatEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *HeartbeatEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "heartbeat",
		Name:        "心跳律动",
		Description: "模拟心脏跳动的律动效果",
		LongDescription: `
心跳律动特效在终端中渲染一个跳动的心形，随着心跳节奏进行缩放动画。

特点：
- ASCII艺术心形图案
- 基于BPM的真实心跳节奏
- 流畅的缩放动画
- 红色脉冲颜色变化
- 自动居中显示

完美用于：
- 浪漫氛围展示
- 健康主题应用
- 节奏可视化
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"动画", "简单", "浪漫", "节奏"},
	}
}

// Init 初始化特效
func (e *HeartbeatEffect) Init(screen tcell.Screen) error {
	e.heartbeat = New(screen, e.config)
	return e.heartbeat.Init()
}

// Run 运行特效
func (e *HeartbeatEffect) Run(quit <-chan struct{}) error {
	return e.heartbeat.Run(quit)
}

// Cleanup 清理资源
func (e *HeartbeatEffect) Cleanup() error {
	if e.heartbeat != nil {
		return e.heartbeat.Cleanup()
	}
	return nil
}
