package oceanwave

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册字符海浪特效到全局注册表
	effects.Register(NewEffect)
}
