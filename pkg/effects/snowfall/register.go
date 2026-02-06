package snowfall

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册雪花飘落特效到全局注册表
	effects.Register(NewEffect)
}
