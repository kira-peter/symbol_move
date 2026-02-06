package digitalwaterfall

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册数字瀑布特效到全局注册表
	effects.Register(NewEffect)
}
