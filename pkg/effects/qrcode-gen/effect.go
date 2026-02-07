package qrcodegen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/symbolmove/symbol_move/pkg/effects"
)

// QRCodeGenEffect 二维码动画特效
type QRCodeGenEffect struct {
	qrcode *QRCodeGen
	config *Config
}

// NewEffect 创建二维码动画特效实例
func NewEffect() effects.Effect {
	return &QRCodeGenEffect{
		config: DefaultConfig(),
	}
}

// Metadata 返回特效元数据
func (e *QRCodeGenEffect) Metadata() effects.Metadata {
	return effects.Metadata{
		ID:            "qrcode-gen",
		Name:          "二维码动画",
		Description:   "生成并展示二维码动画",
		NameEN:        "QR Code Animation",
		DescriptionEN: "Generates and displays animated QR codes",
		LongDescription: `
二维码动画特效生成并展示不同内容的二维码。

特点：
- 使用go-qrcode库生成二维码
- 定期切换不同内容
- 自动缩放以适应屏幕
- 黑白模块清晰渲染
- 显示当前内容文本

完美用于：
- 信息分享
- 链接展示
- 创意二维码应用
`,
		Author:  "SymbolMove",
		Version: "1.0.0",
		Tags:    []string{"实用", "二维码", "动画", "信息"},
	}
}

// Init 初始化特效
func (e *QRCodeGenEffect) Init(screen tcell.Screen) error {
	e.qrcode = New(screen, e.config)
	return e.qrcode.Init()
}

// Run 运行特效
func (e *QRCodeGenEffect) Run(quit <-chan struct{}) error {
	return e.qrcode.Run(quit)
}

// Cleanup 清理资源
func (e *QRCodeGenEffect) Cleanup() error {
	if e.qrcode != nil {
		return e.qrcode.Cleanup()
	}
	return nil
}
