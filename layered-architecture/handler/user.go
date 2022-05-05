package handler

import (
	"net/http"

	"github.com/mmmommm/go-sample/layered-architecture/domain"
	"github.com/mmmommm/go-sample/layered-architecture/usecase"
	"github.com/labstack/echo/v4"
)

type UserGetHandler func(ctx echo.Context) error

type UserGetParams struct{}

func (p *UserGetParams) toInput() *usecase.UserGetCaseInput {
	// TODO:
	return &usecase.UserGetCaseInput{}
}

type UserGetView struct {
	User *domain.User `json:"user"`
}

func applyUserGetView(user *domain.User) *UserGetView {
	return &UserGetView{
		User: user,
	}
}
func ProvideUserGetHandler(uc usecase.UserGetCase) UserGetHandler {
	return func(ctx echo.Context) error {
		// req paramsの解析
		params := &UserGetParams{}
		if err := ctx.Bind(params); err != nil {
			// TODO: logging
			return err
		}
		// usecaseからどんなデータ返す
		result, err := uc(ctx.Request().Context(), params.toInput())
		if err != nil {
			return err
		}

		// usecaseを呼ぶ
		return ctx.JSON(http.StatusOK, applyUserGetView(result))
	}
}