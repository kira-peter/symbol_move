package typewritercode

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册打字机代码雨特效到全局注册表
	effects.Register(NewEffect)
}
