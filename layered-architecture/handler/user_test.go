package handler

import (
	"context"
	"testing"

	"github.com/mmmommm/go-sample/layered-architecture/domain"
	"github.com/mmmommm/go-sample/layered-architecture/usecase"
	"github.com/labstack/echo/v4"
)

func TestUserGetHandler(t *testing.T) {
	testCase := []struct {
		mockGetCase usecase.UserGetCase
		mockCtx     echo.Context
		expect      error
	}{
		{
			mockGetCase: func(ctx context.Context, in *usecase.UserGetCaseInput) (*domain.User, error) {
				return nil, nil
			},
		},
	}

	for _, v := range testCase {
		// usecase層のmock
		handler := ProvideUserGetHandler(v.mockGetCase)
		if err := handler(v.mockCtx); err != v.expect {

		}
	}
}