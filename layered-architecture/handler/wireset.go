package handler

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideUserGetHandler,
	// ProvideHandler,
)