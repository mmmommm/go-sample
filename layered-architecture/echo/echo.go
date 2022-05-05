package echo

import (
	"github.com/mmmommm/go-sample/layered-architecture/handler"
	"github.com/labstack/echo/v4"
)

type EchoServer = *echo.Echo

func ProvideEchoServer(h *handler.Handler) EchoServer {
	e := echo.New()
	// user
	e.GET("/user", echo.HandlerFunc(h.UserGetHandler))

	return e
}