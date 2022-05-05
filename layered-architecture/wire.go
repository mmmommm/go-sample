//go:build wireinject
// +build wireinject

package golang_example

import (
	"github.com/mmmommm/go-sample/layered-architecture/echo"
	"github.com/mmmommm/go-sample/layered-architecture/handler"
	"github.com/mmmommm/go-sample/layered-architecture/repository"
	"github.com/mmmommm/go-sample/layered-architecture/usecase"
	"github.com/google/wire"
)

type EntryPoint struct {
	Srv echo.EchoServer
}

func NewEntryPoint() (*EntryPoint, func(), error) {
	wire.Build(
		echo.Set,
		handler.Set,
		usecase.Set,
		repository.Set,
		wire.Struct(new(EntryPoint), "*"),
	)
	return nil, nil, nil
}