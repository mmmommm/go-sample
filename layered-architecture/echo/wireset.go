package echo

import (
	"github.com/mmmommm/go-sample/layered-architecture/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	handler.ProvideHandler,
	ProvideEchoServer,
)