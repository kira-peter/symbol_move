package waterripple

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册水波涟漪特效到全局注册表
	effects.Register(NewEffect)
}
