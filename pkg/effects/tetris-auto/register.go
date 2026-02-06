package tetrisauto

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册俄罗斯方块AI特效到全局注册表
	effects.Register(NewEffect)
}
