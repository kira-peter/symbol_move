package mazegenerator

import "github.com/symbolmove/symbol_move/pkg/effects"

func init() {
	effects.Register(NewEffect)
}
