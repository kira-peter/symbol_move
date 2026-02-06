package qrcodegen

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册二维码动画特效到全局注册表
	effects.Register(NewEffect)
}
