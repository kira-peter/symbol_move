package heartbeat

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	// 自动注册心跳律动特效到全局注册表
	effects.Register(NewEffect)
}
