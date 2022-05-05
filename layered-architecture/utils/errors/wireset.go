package errors

import "github.com/google/wire"

var Set = wire.NewSet(
	ProvideAppErrorHandler,
)
