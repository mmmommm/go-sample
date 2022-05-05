package errors

import (
	"errors"
	"net/http"

	"github.com/mmmommm/go-sample/logger"
	"github.com/labstack/echo/v4"
)

type ErrorHandler = echo.HTTPErrorHandler
type ErrorView struct {
	Message        string `json:"message"`
	AppErrorID     int    `json:"app_error_id"`
	HttpStatusCode int    `json:"http_status_code"`
}

func newErrorView(appErr *AppError) *ErrorView {
	view := &ErrorView{
		AppErrorID:     appErr.AppErrID,
		HttpStatusCode: appErr.HttpStatusCode,
	}
	switch appErr.ErrorKind {
	case ClientError:
		view.Message = appErr.ClientErrMessage
	case ServerError:
		view.Message = appErr.ServerErrMessage
	case UnknownError:
		view.Message = appErr.ErrMessage
	}
	return view
}

func ProvideAppErrorHandler(log *logger.Logger) ErrorHandler {
	return func(err error, c echo.Context) {
		var (
			code = http.StatusInternalServerError
		)
		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				err := c.NoContent(code)
				if err != nil {
					c.Logger().Error(err)
				}
			} else {
				var appErr *AppError
				if ok := errors.As(err, &appErr); ok {
					// App定義のErrorだった場合
					switch appErr.ErrorKind {
					case ClientError:
						// ClientのErrorはInfoでLogに残す
						log.Infow("Error",
							"kind", ErrorMap[appErr.ErrorKind],
							"error msg", appErr.ClientErrMessage,
							"status code", appErr.HttpStatusCode,
						)
						if err := c.JSON(code, newErrorView(appErr)); err != nil {
							log.Errorw("server error",
								"msg: ", err.Error(),
							)
						}
					case ServerError:
						// Server起因のErrorはInfoでLogに残す
						log.Errorw("Server Error",
							"kind", ErrorMap[appErr.ErrorKind],
							"error msg", appErr.ServerErrMessage,
							"status code", appErr.HttpStatusCode,
						)
						if err := c.JSON(code, newErrorView(appErr)); err != nil {
							log.Errorw("server error",
								"msg: ", err.Error(),
							)
						}
					}
				} else {
					// ひとまずInfoでLogに残す
					// 今は手抜きでサーバーのエラーをそのままclientにも返しちゃう
					appErr := NewAppError(UnknownError, 0, err.Error())
					log.Errorw("Error: ",
						"error msg", appErr.ErrMessage,
					)
					if err := c.JSON(code, newErrorView(appErr)); err != nil {
						log.Errorw("server error",
							"msg: ", appErr.ErrMessage,
						)
					}
				}
			}
		}
	}
}
