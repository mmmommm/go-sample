package usecase

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideUserGetCase,
)