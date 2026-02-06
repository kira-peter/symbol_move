package dnahelix

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// DNAHelixEffect DNA双螺旋特效
type DNAHelixEffect struct {
	dna    *DNAHelix
	config *Config
}

// NewEffect 创建DNA双螺旋特效实例
func NewEffect() effects.Effect {
	return &DNAHelixEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *DNAHelixEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:          "dna-helix",
		Name:        "DNA双螺旋",
		Description: "模拟DNA双螺旋结构的3D旋转动画",
		LongDescription: `
DNA双螺旋特效展示了生命的基本结构 - DNA的双螺旋形态。

特点：
- 3D螺旋数学计算
- 双螺旋链的投影渲染
- 碱基对连接线
- 流畅的旋转动画
- 深度感的颜色变化

完美用于：
- 生物科学演示
- 教育展示
- 科技主题背景
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"3D", "科学", "动画", "教育"},
	}
}

// Init 初始化特效
func (e *DNAHelixEffect) Init(screen tcell.Screen) error {
	e.dna = New(screen, e.config)
	return e.dna.Init()
}

// Run 运行特效
func (e *DNAHelixEffect) Run(quit <-chan struct{}) error {
	return e.dna.Run(quit)
}

// Cleanup 清理资源
func (e *DNAHelixEffect) Cleanup() error {
	if e.dna != nil {
		return e.dna.Cleanup()
	}
	return nil
}
