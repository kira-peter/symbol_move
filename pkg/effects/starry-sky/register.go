package starrysky

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册星空闪烁特效到全局注册表
	effects.Register(NewEffect)
}
