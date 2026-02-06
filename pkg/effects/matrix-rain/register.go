package matrixrain

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册矩阵字符雨特效到全局注册表
	effects.Register(NewEffect)
}
