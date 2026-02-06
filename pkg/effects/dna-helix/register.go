package dnahelix

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册DNA双螺旋特效到全局注册表
	effects.Register(NewEffect)
}
