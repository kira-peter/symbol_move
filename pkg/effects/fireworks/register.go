package fireworks

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册烟花绽放特效到全局注册表
	effects.Register(NewEffect)
}
