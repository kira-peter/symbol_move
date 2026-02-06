package wavetext

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册波浪文字特效到全局注册表
	effects.Register(NewEffect)
}
